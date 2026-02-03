package http_tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"traceforge/api/internal/router"
)

func TestHealth(t *testing.T) {
	r := chi.NewRouter()
	router.Mount(r)

	req := httptest.NewRequest(http.MethodGet, "/api/traceforge/v1/health", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Contains(t, rr.Body.String(), `"ok":true`)
}
