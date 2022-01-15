package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation := '\u2757'
	search := '\U0001F50D'
	weather := '\u2600'

	if strings.ContainsRune(log, recommendation) {
		return "recommendation"
	}

	if strings.ContainsRune(log, search) {
		return "search"
	}

	if strings.ContainsRune(log, weather) {
		return "weather"
	}

	return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
