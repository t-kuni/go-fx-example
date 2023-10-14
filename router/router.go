package router

import (
	"github.com/t-kuni/go-fx-example/handlers"
	"net/http"
)

func NewServeMux(echo *handlers.EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
