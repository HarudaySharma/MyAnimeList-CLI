package utils

import (
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
)

func MapIndicesToDetailFields(detailsIdx []int) []es.AnimeDetailField {
	detailsStr := make([]es.AnimeDetailField, 0)
	for _, v := range detailsIdx {
		if v >= 0 && v < len(es.EveryDetailField()) {
			detailsStr = append(detailsStr, es.EveryDetailField()[v])
		}
	}

	return detailsStr
}
