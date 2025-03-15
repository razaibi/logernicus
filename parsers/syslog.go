package parsers

import (
	"strings"

	"github.com/razaibi/logernicus/models"
)

func ParseSyslog(line string) models.LogEntry {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) < 2 {
		return models.LogEntry{}
	}

	return models.LogEntry{
		Message: parts[1],
	}
}
