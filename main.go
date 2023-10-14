package main

import (
	"github.com/t-kuni/go-fx-example/container"
	"github.com/t-kuni/go-fx-example/services"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	cont := container.NewContainer(
		fx.Decorate(services.NewDummyServiceImplB),
		fx.Invoke(func(server *http.Server) {}),
	)
	cont.Run()
}
