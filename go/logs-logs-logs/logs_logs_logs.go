package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var app rune
	for _, char := range log {
		if char == '❗' || char == '🔍' || char == '☀' {
			app = char
			break
		}

	}
	switch app {
	case '❗':
		return "recommendation"
	case '🔍':
		return "search"
	case '☀':
		return "weather"
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	truc := strings.ReplaceAll(log, string(oldRune), string(newRune))

	return truc
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	nbreChar := utf8.RuneCountInString(log)
	if nbreChar <= limit {
		return true
	} else {
		return false
	}
}
