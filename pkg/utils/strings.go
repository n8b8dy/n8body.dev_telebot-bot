package utils

import "strings"

func JoinMultiline(lines ...string) string {
	return strings.Join(lines, "\n")
}
