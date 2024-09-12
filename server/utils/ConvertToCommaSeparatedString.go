package utils

import (
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/server/enums"
)

func ConvertToCommaSeperatedString(data []enums.AnimeDetailField) string {
	fieldsStr := strings.Builder{}
	for i, d := range data {
		if i != 0 {
			fieldsStr.WriteString(",")
		}
		fieldsStr.WriteString(string(d))
	}

	return fieldsStr.String()
}
