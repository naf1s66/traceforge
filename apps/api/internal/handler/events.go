package handler

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"traceforge/api/internal/model"
)

type eventCreateRequest struct {
	WorkspaceID string         `json:"workspace_id"`
	Service     string         `json:"service"`
	Actor       string         `json:"actor"`
	Action      string         `json:"action"`
	Resource    string         `json:"resource"`
	Metadata    map[string]any `json:"metadata"`
}

type EventStore struct {
	mu     sync.Mutex
	events map[string][]model.Event
	idem   map[string]model.Event
}

func NewEventStore() *EventStore {
	return &EventStore{
		events: make(map[string][]model.Event),
		idem:   make(map[string]model.Event),
	}
}

func (s *EventStore) add(event model.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.WorkspaceID] = append(s.events[event.WorkspaceID], event)
}

func (s *EventStore) getByIdempotency(workspaceID, key string) (model.Event, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	event, ok := s.idem[idempotencyKey(workspaceID, key)]
	return event, ok
}

func (s *EventStore) addWithIdempotency(workspaceID, key string, event model.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.WorkspaceID] = append(s.events[event.WorkspaceID], event)
	s.idem[idempotencyKey(workspaceID, key)] = event
}

func HandleCreateEvent(store *EventStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := strings.TrimSpace(r.Header.Get("X-API-Key"))
		if apiKey == "" {
			writeJSONError(w, http.StatusUnauthorized, "missing_api_key", "X-API-Key is required")
			return
		}
		idempotencyKey := strings.TrimSpace(r.Header.Get("Idempotency-Key"))
		if idempotencyKey == "" {
			writeJSONError(w, http.StatusBadRequest, "missing_idempotency_key", "Idempotency-Key is required")
			return
		}
		expectedKey, err := writeAPIKey()
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "write_api_key_unset", err.Error())
			return
		}
		if apiKey != expectedKey {
			writeJSONError(w, http.StatusForbidden, "invalid_api_key", "X-API-Key is invalid")
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		var req eventCreateRequest
		if err := decoder.Decode(&req); err != nil {
			writeJSONError(w, http.StatusBadRequest, "invalid_json", err.Error())
			return
		}
		if err := ensureEOF(decoder); err != nil {
			writeJSONError(w, http.StatusBadRequest, "invalid_json", "request body must be a single JSON object")
			return
		}

		if err := validateCreateEvent(req); err != nil {
			writeJSONError(w, http.StatusBadRequest, "invalid_request", err.Error())
			return
		}

		workspaceID := strings.TrimSpace(req.WorkspaceID)
		if existing, ok := store.getByIdempotency(workspaceID, idempotencyKey); ok {
			writeJSON(w, http.StatusOK, existing)
			return
		}

		metadata := req.Metadata
		if metadata == nil {
			metadata = map[string]any{}
		}

		eventID, err := newUUID()
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "id_generation_failed", "could not generate event id")
			return
		}

		event := model.Event{
			ID:          eventID,
			WorkspaceID: workspaceID,
			Service:     strings.TrimSpace(req.Service),
			Actor:       strings.TrimSpace(req.Actor),
			Action:      strings.TrimSpace(req.Action),
			Resource:    strings.TrimSpace(req.Resource),
			Metadata:    metadata,
			CreatedAt:   time.Now().UTC(),
		}

		store.addWithIdempotency(workspaceID, idempotencyKey, event)

		writeJSON(w, http.StatusCreated, event)
	}
}

func validateCreateEvent(req eventCreateRequest) error {
	if strings.TrimSpace(req.WorkspaceID) == "" {
		return errRequiredField("workspace_id")
	}
	if strings.TrimSpace(req.Service) == "" {
		return errRequiredField("service")
	}
	if strings.TrimSpace(req.Actor) == "" {
		return errRequiredField("actor")
	}
	if strings.TrimSpace(req.Action) == "" {
		return errRequiredField("action")
	}
	if strings.TrimSpace(req.Resource) == "" {
		return errRequiredField("resource")
	}
	return nil
}

type apiError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func writeJSONError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, apiError{
		Error:   code,
		Message: message,
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload == nil {
		return
	}
	_ = json.NewEncoder(w).Encode(payload)
}

func ensureEOF(decoder *json.Decoder) error {
	if err := decoder.Decode(&struct{}{}); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	return io.ErrUnexpectedEOF
}

func errRequiredField(field string) error {
	return &requestFieldError{Field: field}
}

type requestFieldError struct {
	Field string
}

func (e *requestFieldError) Error() string {
	return "missing required field: " + e.Field
}

func writeAPIKey() (string, error) {
	if key := strings.TrimSpace(os.Getenv("WRITE_API_KEY")); key != "" {
		return key, nil
	}
	return "", fmt.Errorf("WRITE_API_KEY is not configured")
}

func newUUID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}

func idempotencyKey(workspaceID, key string) string {
	return workspaceID + ":" + key
}
