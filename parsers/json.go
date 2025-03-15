package parsers

import (
	"encoding/json"

	"github.com/razaibi/logernicus/models"
)

func ParseJSON(line string) models.LogEntry {
	var entry models.LogEntry
	json.Unmarshal([]byte(line), &entry)
	return entry
}
