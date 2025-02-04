package proxy

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omaqase/sato/gateway/config"
	userService "github.com/omaqase/sato/gateway/pkg/api/v1/user"
	"github.com/omaqase/sato/gateway/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
	"net/http"
)

type Proxy struct {
	Mux    *runtime.ServeMux
	Config config.Config
}

func NewProxy(config config.Config) *Proxy {
	return &Proxy{Mux: runtime.NewServeMux(), Config: config}
}

func (p *Proxy) RegisterServices(ctx context.Context) error {
	if err := p.registerUserService(ctx); err != nil {
		return fmt.Errorf("failed to register user service: %w", err)
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
			"/api/v1/auth/sign-in":  true,
			"/api/v	1/auth/sign-up": true,
		}

		if publicPaths[r.URL.Path] {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("dqwew")
			log.Println(authHeader)
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := utils.ParseJWT("qxwneiqwheuiys", authHeader)
		if err != nil {
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Grpc-Metadata-X-Request-ID")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
