package utils

import "strings"

func JoinMultiline(lines ...string) string {
	return strings.Join(lines, "\n")
}

func EscapeHTML(str string) string {
	escapeChars := map[string]string{
		"&": "&amp;",
		"<": "&lt;",
		">": "&gt;",
	}

	for key, val := range escapeChars {
		str = strings.ReplaceAll(str, key, val)
	}

	return str
}
