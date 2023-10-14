package server

import (
	"context"
	"github.com/t-kuni/go-fx-example/logger"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func NewHTTPServer(
	lc fx.Lifecycle,
	mux *http.ServeMux,
	log *logger.Logger,
) *http.Server {
	srv := &http.Server{Addr: ":44444", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.SimpleInfoF("Starting HTTP server")

			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
