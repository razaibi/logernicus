package logernicus

import (
	"encoding/json"
	"regexp"
	"strings"
)

// Detect log format based on a single log line
func detectFormat(line string) string {
	if isJSON(line) {
		return "json"
	} else if isCLF(line) {
		return "clf"
	} else if isKV(line) {
		return "kv"
	} else if isSyslog(line) {
		return "syslog"
	} else if isApache(line) {
		return "apache"
	}
	return "unknown"
}

// JSON detection
func isJSON(line string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(line), &js) == nil
}

// Common Log Format detection
func isCLF(line string) bool {
	clfRegex := `^(\S+) - (\S+) \[(.*?)\] "(\S+) (.*?) (\S+)" (\d+) (\d+|-)$`
	matched, _ := regexp.MatchString(clfRegex, line)
	return matched
}

// Key-Value Format detection
func isKV(line string) bool {
	return strings.Contains(line, "=") && strings.Contains(line, " ")
}

// Syslog detection
func isSyslog(line string) bool {
	syslogRegex := `^<\d+>\d+`
	matched, _ := regexp.MatchString(syslogRegex, line)
	return matched
}

// Apache Log detection
func isApache(line string) bool {
	apacheRegex := `^(\S+) - (\S+) \[(.*?)\] "(\S+) (.*?) (\S+)" (\d+) (\d+) "(.*?)" "(.*?)"$`
	matched, _ := regexp.MatchString(apacheRegex, line)
	return matched
}
