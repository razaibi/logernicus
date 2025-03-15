package parsers

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/razaibi/logernicus"
)

// Regular expression to match Apache/Nginx access logs
var apacheRegex = regexp.MustCompile(`^(\S+) - (\S+) \[(.*?)\] "(\S+) (.*?) (\S+)" (\d+) (\d+) "(.*?)" "(.*?)"$`)

// ParseApache parses an Apache/Nginx access log line into a LogEntry struct
func ParseApache(line string) logernicus.LogEntry {
	matches := apacheRegex.FindStringSubmatch(line)
	if matches == nil {
		return logernicus.LogEntry{}
	}

	// Convert status code and bytes sent to integers
	status, _ := strconv.Atoi(matches[7])
	bytesSent, _ := strconv.Atoi(matches[8])

	return logernicus.LogEntry{
		IP:         matches[1],
		Timestamp:  matches[3],
		Request:    fmt.Sprintf("%s %s", matches[4], matches[5]),
		StatusCode: status,
		Message:    fmt.Sprintf("User Agent: %s, Referrer: %s, Bytes Sent: %d", matches[10], matches[9], bytesSent),
	}
}
