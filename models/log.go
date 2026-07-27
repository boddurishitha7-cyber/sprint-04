package models

import "time"

type Log struct {
	TraceID    string   
	ID          int
	ServiceName string
	LogLevel    string
	Message      string
	FailureType string
	EventTime   time.Time
}