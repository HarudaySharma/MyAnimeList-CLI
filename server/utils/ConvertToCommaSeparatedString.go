package utils

import (
	"strings"
)

func ConvertToCommaSeperatedString(data []AnimeDetailField) string {
	fieldsStr := strings.Builder{}
	for i, d := range data {
		if i != 0 {
			fieldsStr.WriteString(",")
		}
		fieldsStr.WriteString(string(d))
	}

	return fieldsStr.String()
}
