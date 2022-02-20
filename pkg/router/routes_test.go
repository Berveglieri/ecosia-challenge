package router

import (
	"ecosia/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func router() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/tree", handlers.TreeHandler)
	router.Get("/tree/{name}", handlers.TreeHandlerWithParams)
	router.Get("/status", handlers.StatusHandler)

	return router
}

func TestTreeEndpoint(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/tree", nil)
	response := httptest.NewRecorder()
	router().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "200 status code is expected")

}

func TestTreeEndpointWithParams(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/tree/Cajueiro", nil)
	response := httptest.NewRecorder()
	router().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "200 status code is expected")
}

func TestStatus(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/status", nil)
	response := httptest.NewRecorder()
	router().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "200 status is expected")
}
