package parsers

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/razaibi/logernicus/models"
)

var clfRegex = regexp.MustCompile(`^(\S+) - (\S+) \[(.*?)\] "(\S+) (.*?) (\S+)" (\d+) (\d+|-)$`)

func ParseCLF(line string) models.LogEntry {
	matches := clfRegex.FindStringSubmatch(line)
	if matches == nil {
		return models.LogEntry{}
	}

	status, _ := strconv.Atoi(matches[7])

	return models.LogEntry{
		IP:         matches[1],
		Timestamp:  matches[3],
		Request:    fmt.Sprintf("%s %s", matches[4], matches[5]),
		StatusCode: status,
	}
}
