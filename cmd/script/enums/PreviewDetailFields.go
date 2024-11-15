package enums

import es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"

var previewDetailFields []es.AnimeDetailField

func PreviewDetailFields() *[]es.AnimeDetailField {
	return &previewDetailFields
}

func init() {
	previewDetailFields = []es.AnimeDetailField{
		es.MainPicture,
		es.AlternativeTitles,
        es.Genres,
		es.Mean, es.Status, es.Broadcast,
		es.NumEpisodes, es.AverageEpisodeDuration, es.EndDate, es.StartDate,
	}
}
