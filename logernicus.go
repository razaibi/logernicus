package logernicus

import (
	"bufio"
	"fmt"
	"os"

	"github.com/razaibi/logernicus/parsers"
)

// LogEntry represents a generic log entry
type LogEntry struct {
	Timestamp  string
	Level      string
	Message    string
	IP         string
	UserAgent  string
	Request    string
	StatusCode int
}

// ReadLogFile automatically detects and parses a log file
func ReadLogFile(filename string) ([]LogEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []LogEntry
	for scanner.Scan() {
		line := scanner.Text()

		// Identify log format
		format := detectFormat(line)

		var entry LogEntry
		switch format {
		case "clf":
			entry = parsers.ParseCLF(line)
		case "json":
			entry = parsers.ParseJSON(line)
		case "kv":
			entry = parsers.ParseKV(line)
		case "syslog":
			entry = parsers.ParseSyslog(line)
		case "apache":
			entry = parsers.ParseApache(line)
		default:
			fmt.Println("Unknown format: ", line)
			continue
		}

		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
