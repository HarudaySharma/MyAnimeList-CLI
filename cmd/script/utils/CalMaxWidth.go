package utils

import "strings"

func CalMaxWidth(text string) int {
	rows := strings.Split(text, "\n")
	maxLen := 0
	for _, r := range rows {
		maxLen = max(maxLen, len(r))
	}

	return maxLen
}

