package main

import (
	"context"
	"github.com/t-kuni/go-fx-example/a/a"
	a2 "github.com/t-kuni/go-fx-example/b/a"
	"github.com/t-kuni/go-fx-example/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"io"
	"net"
	"net/http"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *logger.Logger) fxevent.Logger {
			return log
		}),
		fx.Provide(
			NewHTTPServer,
			NewServeMux,
			NewEchoHandler,
			a.NewTest,
			a2.NewTest,
			logger.NewLogger,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *logger.Logger, t1 *a.Test, t2 *a2.Test) *http.Server {
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
	h.log.SimpleInfoF("Hello from EchoHandler")
}
func NewServeMux(echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
