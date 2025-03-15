package parsers

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/razaibi/logernicus"
)

var clfRegex = regexp.MustCompile(`^(\S+) - (\S+) \[(.*?)\] "(\S+) (.*?) (\S+)" (\d+) (\d+|-)$`)

func ParseCLF(line string) logernicus.LogEntry {
	matches := clfRegex.FindStringSubmatch(line)
	if matches == nil {
		return logernicus.LogEntry{}
	}

	status, _ := strconv.Atoi(matches[7])

	return logernicus.LogEntry{
		IP:         matches[1],
		Timestamp:  matches[3],
		Request:    fmt.Sprintf("%s %s", matches[4], matches[5]),
		StatusCode: status,
	}
}
