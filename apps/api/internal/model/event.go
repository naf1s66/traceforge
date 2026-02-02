package model

import "time"

type Event struct {
	ID          string         `json:"id"`
	WorkspaceID string         `json:"workspace_id"`
	Service     string         `json:"service"`
	Actor       string         `json:"actor"`
	Action      string         `json:"action"`
	Resource    string         `json:"resource"`
	Metadata    map[string]any `json:"metadata"`
	CreatedAt   time.Time      `json:"created_at"`
}
