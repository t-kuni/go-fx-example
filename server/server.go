package server

import (
	"context"
	a1 "github.com/t-kuni/go-fx-example/a/a"
	a2 "github.com/t-kuni/go-fx-example/b/a"
	"github.com/t-kuni/go-fx-example/logger"
	"github.com/t-kuni/go-fx-example/services"
	"go.uber.org/fx"
	"io"
	"net"
	"net/http"
)

func NewHTTPServer(
	lc fx.Lifecycle,
	mux *http.ServeMux,
	log *logger.Logger,
	t1 *a1.Test,
	t2 *a2.Test,
	s services.IDummyService,
) *http.Server {
	srv := &http.Server{Addr: ":44444", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.SimpleInfoF("Starting HTTP server")

			t1.Hello()
			t2.Hello()
			s.Hello()

			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

type EchoHandler struct {
	log *logger.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *logger.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.SimpleFatal(err, nil)
	}
	h.log.SimpleInfoF("hello from echo handler")
}
func NewServeMux(echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
