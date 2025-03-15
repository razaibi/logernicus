package logernicus

import (
	"bufio"
	"fmt"
	"os"

	"logernicus/models"
	"logernicus/parsers"
)

// ReadLogFile automatically detects and parses a log file
func ReadLogFile(filename string) ([]models.LogEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []models.LogEntry
	for scanner.Scan() {
		line := scanner.Text()

		// Identify log format
		format := detectFormat(line)

		var entry models.LogEntry
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
