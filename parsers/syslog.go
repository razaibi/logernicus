package parsers

import (
	"strings"

	"github.com/razaibi/logernicus"
)

func ParseSyslog(line string) logernicus.LogEntry {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) < 2 {
		return logernicus.LogEntry{}
	}

	return logernicus.LogEntry{
		Message: parts[1],
	}
}
