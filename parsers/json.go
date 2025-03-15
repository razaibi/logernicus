package parsers

import (
	"encoding/json"

	"logernicus/models"
)

func ParseJSON(line string) models.LogEntry {
	var entry models.LogEntry
	json.Unmarshal([]byte(line), &entry)
	return entry
}
