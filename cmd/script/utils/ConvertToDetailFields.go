package utils

import (
	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
)

func ConvertToDetailFields(detailsIdx []int) ([]es.AnimeDetailField) {
	detailsStr := make([]string, 0)
	for _, v := range detailsIdx {
		if v < len(e.DetailsFields) {
			detailsStr = append(detailsStr, e.DetailsFields[v])
		}
	}

    details, _ := es.ParseDetailsField(detailsStr);
    return details
}
