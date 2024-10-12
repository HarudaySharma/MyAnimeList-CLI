package utils

import "regexp"

// Helper function to strip ANSI codes from a string
func StripAnsi(str string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(str, "")
}
