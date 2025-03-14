package proxy

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"slices"

	protobuf "github.com/omaqase/sato/gateway/pkg/api/v1/catalogue"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omaqase/sato/gateway/config"
	notificationService "github.com/omaqase/sato/gateway/pkg/api/v1/notification"
	userService "github.com/omaqase/sato/gateway/pkg/api/v1/user"
	"github.com/omaqase/sato/gateway/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/proto"
)

type Proxy struct {
	Mux    *runtime.ServeMux
	Config config.Config
}

// Новый конструктор с дополнительной опцией для обработки ServerMetadata.
func NewProxy(config config.Config) *Proxy {
	// Используем опцию WithForwardResponseOption, чтобы обработать ситуацию,
	// когда ServerMetadata отсутствует в контексте.
	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
			// Пытаемся извлечь ServerMetadata. Если его нет, ничего не делаем.
			_, ok := runtime.ServerMetadataFromContext(ctx)
			if !ok {
				// Можно залогировать, что metadata не найдены (если нужно).
				grpclog.Warning("ServerMetadata not found in context")
			}
			return nil
		}),
	)
	return &Proxy{Mux: mux, Config: config}
}

func (p *Proxy) RegisterServices(ctx context.Context) error {
	if err := p.registerUserService(ctx); err != nil {
		return fmt.Errorf("failed to register user service: %w", err)
	}

	if err := p.registerNotificationService(ctx); err != nil {
		return fmt.Errorf("failed to register notification service: %w", err)
	}

	if err := p.registerProductService(ctx); err != nil {
		return fmt.Errorf("failed to register product service: %w", err)
	}

	return nil
}

func (p *Proxy) registerUserService(ctx context.Context) error {
	endpoint := fmt.Sprintf("%s:%s", p.Config.UserService.Host, p.Config.UserService.Port)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := userService.RegisterUserServiceHandlerFromEndpoint(ctx, p.Mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}

func (p *Proxy) registerNotificationService(ctx context.Context) error {
	endpoint := fmt.Sprintf("%s:%s", p.Config.NotificationService.Host, p.Config.NotificationService.Port)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := notificationService.RegisterNotificationServiceHandlerFromEndpoint(ctx, p.Mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}

func (p *Proxy) registerProductService(ctx context.Context) error {
	endpoint := fmt.Sprintf("%s:%s", p.Config.ProductService.Host, p.Config.ProductService.Port)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := protobuf.RegisterCategoryServiceHandlerFromEndpoint(ctx, p.Mux, endpoint, opts); err != nil {
		return err
	}

	if err := protobuf.RegisterProductServiceHandlerFromEndpoint(ctx, p.Mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}

func (p *Proxy) Serve() error {
	handler := EnableCORS(AuthenticationMiddleware(p.Mux))
	addr := fmt.Sprintf(":%s", p.Config.App.Port)

	grpclog.Infof("Starting gRPC Gateway server on %s", addr)
	return http.ListenAndServe(addr, handler)
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		publicPaths := map[string]bool{
			"/api/v1/auth/send-otp":   true,
			"/api/v1/auth/verify-otp": true,
		}

		if publicPaths[r.URL.Path] {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println(3)
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ParseJWT("qxwneiqwheuiys", authHeader)
		if err != nil {
			log.Println(claims)
			log.Println(2)
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}

		roles := GetRequiredRoles(r.URL.Path)

		if len(roles) == 0 || !slices.Contains(roles, claims.Role) {
			log.Println(roles)
			log.Println(claims.Role)
			log.Println(3)
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Grpc-Metadata-X-Request-ID, X-CSRF-Token, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
