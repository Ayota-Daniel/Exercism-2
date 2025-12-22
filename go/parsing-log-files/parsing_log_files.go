package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	r := regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	r := regexp.MustCompile(`<{1}[~\*=-]*>{1}`)
	return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	r := regexp.MustCompile(`"{1}.*(?i)password"{1}`)
	for _, line := range lines {
		if r.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`\w*end-of-line[0-9]*`)
	text = r.ReplaceAllString(text, "")
	return text
}

func TagWithUserName(lines []string) []string {
	r := regexp.MustCompile(`User\s+\w+`)
	for i, line := range lines {
		if temp := r.FindString(line); temp != "" {
			temp = strings.ReplaceAll(temp, "User ", "")
			temp = strings.Trim(temp, " ")
			lines[i] = fmt.Sprintf("[USR] %s %s", temp, line)
		}
	}
	return lines
}
