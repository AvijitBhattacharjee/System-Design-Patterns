package distributedloggingsystem

import "time"

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type LogEntry struct {
	LogID     string
	Timestamp time.Time

	Service string
	Host    string

	Level   LogLevel
	Message string

	TraceID string
	SpanID  string

	Fields map[string]interface{}
}