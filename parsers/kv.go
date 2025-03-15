package parsers

import (
	"strings"

	"github.com/razaibi/logernicus"
)

func ParseKV(line string) logernicus.LogEntry {
	fields := strings.Fields(line)
	entry := logernicus.LogEntry{}

	for _, field := range fields {
		parts := strings.SplitN(field, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]

		switch key {
		case "timestamp":
			entry.Timestamp = value
		case "level":
			entry.Level = value
		case "message":
			entry.Message = value
		case "ip":
			entry.IP = value
		}
	}

	return entry
}
