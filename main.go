package main

import (
	"github.com/t-kuni/go-fx-example/container"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	cont := container.NewContainer(
		fx.Invoke(func(server *http.Server) {}),
	)
	cont.Run()
}
