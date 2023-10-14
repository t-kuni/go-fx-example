package container

import (
	a1 "github.com/t-kuni/go-fx-example/a/a"
	a2 "github.com/t-kuni/go-fx-example/b/a"
	"github.com/t-kuni/go-fx-example/handlers"
	"github.com/t-kuni/go-fx-example/logger"
	"github.com/t-kuni/go-fx-example/router"
	"github.com/t-kuni/go-fx-example/server"
	"github.com/t-kuni/go-fx-example/services"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewContainer(opts ...fx.Option) *fx.App {
	mergedOpts := []fx.Option{
		fx.WithLogger(func(log *logger.Logger) fxevent.Logger {
			return log
		}),
		fx.Provide(
			server.NewHTTPServer,
			router.NewServeMux,
			handlers.NewEchoHandler,
			a1.NewTest,
			a2.NewTest,
			logger.NewLogger,
			services.NewDummyServiceImplA,
		),
	}
	mergedOpts = append(mergedOpts, opts...)

	return fx.New(mergedOpts...)
}
