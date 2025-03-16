package logernicus

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/razaibi/logernicus/models"
)

var (
	// Common timestamp formats found in logs
	timeFormats = []string{
		"02/Jan/2006:15:04:05 -0700", // Apache/Common Log Format
		"Jan 02 15:04:05",            // Syslog
		"2006-01-02T15:04:05Z",       // ISO8601/RFC3339
		"2006-01-02T15:04:05.999Z",   // ISO8601 with milliseconds
		"2006/01/02 15:04:05",        // Common datetime format
		"2006-01-02 15:04:05",        // Another common datetime format
		"02/Jan/2006 15:04:05",       // Alternative Apache format
		"Mon Jan 02 15:04:05 2006",   // Unix date
		"Jan 02 15:04:05 MST 2006",   // Alternative syslog
	}

	// Extract timestamp pattern from a log line
	timestampPatterns = []*regexp.Regexp{
		regexp.MustCompile(`\[([^]]+)\]`),                               // Matches [timestamp]
		regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.\d+)`), // ISO format with ms
		regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2})`),     // ISO format
		regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2})`),     // YYYY/MM/DD HH:MM:SS
		regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`),     // YYYY-MM-DD HH:MM:SS
		regexp.MustCompile(`([A-Z][a-z]{2} \d{2} \d{2}:\d{2}:\d{2})`),   // Syslog format
	}
)

// ExtractTimestamp attempts to extract a timestamp string from a log line
func ExtractTimestamp(logLine string) string {
	for _, pattern := range timestampPatterns {
		matches := pattern.FindStringSubmatch(logLine)
		if len(matches) > 1 {
			return matches[1]
		}
	}
	return ""
}

// DetectTimeFormat tries to determine the timestamp format used in logs
func DetectTimeFormat(entries []models.LogEntry) string {
	if len(entries) == 0 {
		return timeFormats[0] // Default to Common Log Format
	}

	// Try various timestamps from the entries
	for _, entry := range entries {
		if entry.Timestamp == "" {
			continue
		}

		// Try to parse with each format
		for _, format := range timeFormats {
			_, err := time.Parse(format, entry.Timestamp)
			if err == nil {
				return format
			}
		}
	}

	return timeFormats[0] // Default if nothing is detected
}

// EnrichTimestamps attempts to parse and normalize all timestamps in log entries
func EnrichTimestamps(entries []models.LogEntry) []models.LogEntry {
	format := DetectTimeFormat(entries)

	for i, entry := range entries {
		if entry.Timestamp == "" {
			continue
		}

		// Try to parse with the detected format
		t, err := time.Parse(format, entry.Timestamp)
		if err == nil {
			// If successful, normalize to ISO format
			entries[i].Timestamp = t.Format(time.RFC3339)
		}
	}

	return entries
}

// ParseTimeRange creates a time range from a string expression
// Examples: "last 24h", "last 7d", "2023-02-01 to 2023-02-28"
func ParseTimeRange(timeRange string, format string) (time.Time, time.Time, error) {
	var startTime, endTime time.Time
	now := time.Now()

	// Handle relative time ranges
	if strings.HasPrefix(timeRange, "last ") {
		duration := strings.TrimPrefix(timeRange, "last ")
		endTime = now

		// Parse the duration (e.g., "24h", "7d", "30m")
		var value int
		var unit string
		fmt.Sscanf(duration, "%d%s", &value, &unit)

		switch unit {
		case "m", "min":
			startTime = now.Add(time.Duration(-value) * time.Minute)
		case "h", "hr", "hrs":
			startTime = now.Add(time.Duration(-value) * time.Hour)
		case "d", "day", "days":
			startTime = now.AddDate(0, 0, -value)
		case "w", "week", "weeks":
			startTime = now.AddDate(0, 0, -value*7)
		case "M", "month", "months":
			startTime = now.AddDate(0, -value, 0)
		case "y", "year", "years":
			startTime = now.AddDate(-value, 0, 0)
		default:
			return time.Time{}, time.Time{}, fmt.Errorf("invalid time unit: %s", unit)
		}
	} else if strings.Contains(timeRange, " to ") {
		// Handle absolute time range: "start to end"
		parts := strings.Split(timeRange, " to ")
		if len(parts) != 2 {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid time range format")
		}

		var err error
		startTime, err = time.Parse(format, strings.TrimSpace(parts[0]))
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid start time: %v", err)
		}

		endTime, err = time.Parse(format, strings.TrimSpace(parts[1]))
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid end time: %v", err)
		}
	} else {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid time range format")
	}

	return startTime, endTime, nil
}
