package parsers

import (
	"encoding/json"

	"github.com/razaibi/logernicus"
)

func ParseJSON(line string) logernicus.LogEntry {
	var entry logernicus.LogEntry
	json.Unmarshal([]byte(line), &entry)
	return entry
}
