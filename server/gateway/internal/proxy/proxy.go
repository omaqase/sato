package proxy

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omaqase/sato/gateway/config"
)

type Proxy struct {
	Mux *runtime.ServeMux
}

func NewProxy() *Proxy {
	return &Proxy{Mux: runtime.NewServeMux()}
}

func (p *Proxy) RegisterNotificationServiceHandler(ctx context.Context, config config.NotificationServiceConfig) {

}
