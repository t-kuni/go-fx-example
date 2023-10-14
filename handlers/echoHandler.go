package handlers

import (
	a1 "github.com/t-kuni/go-fx-example/a/a"
	a2 "github.com/t-kuni/go-fx-example/b/a"
	"github.com/t-kuni/go-fx-example/logger"
	"github.com/t-kuni/go-fx-example/services"
	"net/http"
)

type EchoHandler struct {
	log *logger.Logger
	t1  *a1.Test
	t2  *a2.Test
	s   services.IDummyService
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(
	log *logger.Logger,
	t1 *a1.Test,
	t2 *a2.Test,
	s services.IDummyService,
) *EchoHandler {
	return &EchoHandler{
		log,
		t1,
		t2,
		s,
	}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.SimpleInfoF("hello from echo handler")
	h.t1.Hello()
	h.t2.Hello()
	println(h.s.Hello())
}
