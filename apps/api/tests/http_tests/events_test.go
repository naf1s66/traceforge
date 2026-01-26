package http_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"traceforge/api/internal/router"
)

func TestCreateEventRequiresAPIKey(t *testing.T) {
	t.Setenv("WRITE_API_KEY", "test-key")

	r := chi.NewRouter()
	router.Mount(r)

	req := httptest.NewRequest(http.MethodPost, "/api/traceforge/v1/events", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestCreateEventRejectsInvalidAPIKey(t *testing.T) {
	t.Setenv("WRITE_API_KEY", "test-key")

	r := chi.NewRouter()
	router.Mount(r)

	body := bytes.NewBufferString(`{"workspace_id":"public","service":"demo","actor":"user:1","action":"LOGIN","resource":"session:1"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/traceforge/v1/events", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "wrong-key")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusForbidden, rr.Code)
}

func TestCreateEventValidatesRequiredFields(t *testing.T) {
	t.Setenv("WRITE_API_KEY", "test-key")

	r := chi.NewRouter()
	router.Mount(r)

	body := bytes.NewBufferString(`{"workspace_id":"public","service":"demo"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/traceforge/v1/events", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "test-key")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateEventSuccess(t *testing.T) {
	t.Setenv("WRITE_API_KEY", "test-key")

	r := chi.NewRouter()
	router.Mount(r)

	body := bytes.NewBufferString(`{"workspace_id":"public","service":"demo","actor":"user:1","action":"LOGIN","resource":"session:1","metadata":{"ip":"127.0.0.1"}}`)
	req := httptest.NewRequest(http.MethodPost, "/api/traceforge/v1/events", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "test-key")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&resp))
	require.NotEmpty(t, resp["id"])
	require.Equal(t, "public", resp["workspace_id"])
	require.Equal(t, "demo", resp["service"])
	require.Equal(t, "user:1", resp["actor"])
	require.Equal(t, "LOGIN", resp["action"])
	require.Equal(t, "session:1", resp["resource"])
	require.NotEmpty(t, resp["created_at"])
}
