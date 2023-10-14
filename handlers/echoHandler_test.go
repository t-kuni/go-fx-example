package handlers_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/t-kuni/go-fx-example/container"
	"github.com/t-kuni/go-fx-example/handlers"
	"github.com/t-kuni/go-fx-example/services"
	"go.uber.org/fx"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler_ServeHTTP(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	t.Run("aaaa", func(t *testing.T) {
		ctx := context.TODO()

		mockDummyService := services.NewMockIDummyService(mockCtrl)
		mockDummyService.EXPECT().Hello().Return("test")

		err := container.NewContainer(
			fx.Decorate(func() services.IDummyService { return mockDummyService }),
			fx.Invoke(func(h *handlers.EchoHandler) {
				req, err := http.NewRequest("GET", "http://example.com/foo", nil)
				assert.NoError(t, err)
				rr := httptest.NewRecorder()
				h.ServeHTTP(rr, req)
			}),
		).Start(ctx)
		assert.NoError(t, err)
	})
}
