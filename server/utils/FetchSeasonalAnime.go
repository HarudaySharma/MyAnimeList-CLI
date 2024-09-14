package utils

import "github.com/HarudaySharma/MyAnimeList-CLI/server/enums"

type FetchSeasonalAnimeParams struct {
	Season string
	Year   string
	Sort   string
	Limit  int
	Offset int
	Fields []enums.AnimeDetailField
}

func FetchSeasonalAnime() {
}
