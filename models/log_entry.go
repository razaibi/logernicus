package models

// LogEntry represents a generic log entry
type LogEntry struct {
	Timestamp  string
	Level      string
	Message    string
	IP         string
	UserAgent  string
	Request    string
	StatusCode int
}
