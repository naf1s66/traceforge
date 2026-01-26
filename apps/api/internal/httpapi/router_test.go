package httpapi

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/go-chi/chi/v5"
  "github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
  r := chi.NewRouter()
  Mount(r)

  req := httptest.NewRequest(http.MethodGet, "/api/traceforge/v1/health", nil)
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  require.Equal(t, http.StatusOK, rr.Code)
  require.Contains(t, rr.Body.String(), `"ok":true`)
}
