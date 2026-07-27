package models

type Health struct {
	TraceID string   `json:"TraceID"`
	Status    string `json:"status"`
	Service   string `json:"service"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}