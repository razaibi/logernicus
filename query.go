package logernicus

import (
	"strings"
	"time"

	"github.com/razaibi/logernicus/models"
)

// QueryOptions defines all possible filtering options for querying logs
type QueryOptions struct {
	IP            string    // Filter by IP address
	StatusCode    int       // Filter by status code
	Level         string    // Filter by log level
	Contains      string    // Filter logs containing this text
	StartTime     time.Time // Filter logs after this timestamp
	EndTime       time.Time // Filter logs before this timestamp
	RequestMethod string    // Filter by HTTP method (GET, POST, etc.)
	MinStatus     int       // Filter by status code >= this value
	MaxStatus     int       // Filter by status code <= this value
	Limit         int       // Limit number of results returned
}

// Query filters log entries based on the provided options
func Query(entries []models.LogEntry, options QueryOptions) []models.LogEntry {
	var results []models.LogEntry

	// Parse the layout based on existing logs
	timeLayout := "02/Jan/2006:15:04:05 -0700" // Default Common Log Format

	for _, entry := range entries {
		// Skip if doesn't match IP filter
		if options.IP != "" && entry.IP != options.IP {
			continue
		}

		// Skip if doesn't match status code filter
		if options.StatusCode != 0 && entry.StatusCode != options.StatusCode {
			continue
		}

		// Skip if doesn't match level filter
		if options.Level != "" && !strings.EqualFold(entry.Level, options.Level) {
			continue
		}

		// Skip if doesn't match contains filter
		if options.Contains != "" &&
			!strings.Contains(strings.ToLower(entry.Message), strings.ToLower(options.Contains)) &&
			!strings.Contains(strings.ToLower(entry.Request), strings.ToLower(options.Contains)) {
			continue
		}

		// Skip if doesn't match request method filter
		if options.RequestMethod != "" &&
			!strings.HasPrefix(strings.ToUpper(entry.Request), strings.ToUpper(options.RequestMethod)+" ") {
			continue
		}

		// Skip if doesn't match min status filter
		if options.MinStatus != 0 && entry.StatusCode < options.MinStatus {
			continue
		}

		// Skip if doesn't match max status filter
		if options.MaxStatus != 0 && entry.StatusCode > options.MaxStatus {
			continue
		}

		// Skip based on time range filters
		if !options.StartTime.IsZero() || !options.EndTime.IsZero() {
			entryTime, err := time.Parse(timeLayout, entry.Timestamp)
			if err == nil {
				if !options.StartTime.IsZero() && entryTime.Before(options.StartTime) {
					continue
				}
				if !options.EndTime.IsZero() && entryTime.After(options.EndTime) {
					continue
				}
			}
		}

		results = append(results, entry)

		// Apply limit if set
		if options.Limit > 0 && len(results) >= options.Limit {
			break
		}
	}

	return results
}

// Count returns the number of log entries matching the given options
func Count(entries []models.LogEntry, options QueryOptions) int {
	matched := Query(entries, options)
	return len(matched)
}

// GroupBy groups log entries by a specific field and returns counts
func GroupBy(entries []models.LogEntry, field string) map[string]int {
	results := make(map[string]int)

	for _, entry := range entries {
		var key string

		switch field {
		case "ip":
			key = entry.IP
		case "level":
			key = entry.Level
		case "status":
			key = string(rune(entry.StatusCode))
		case "request":
			// Extract method from request if possible
			parts := strings.Fields(entry.Request)
			if len(parts) > 0 {
				key = parts[0]
			} else {
				key = entry.Request
			}
		default:
			key = "unknown"
		}

		if key != "" {
			results[key]++
		}
	}

	return results
}
