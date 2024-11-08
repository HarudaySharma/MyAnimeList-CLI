package utils

import "strings"

func CalMaxHeight(text string) int {
	rows := strings.Split(text, "\n")

	return len(rows)
}
