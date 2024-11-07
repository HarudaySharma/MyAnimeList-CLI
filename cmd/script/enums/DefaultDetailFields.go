package enums

import es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"

var defaultDetailFields []es.AnimeDetailField
var defaultDetailFieldsMap map[es.AnimeDetailField]bool

func DefaultDetailFields() *[]es.AnimeDetailField {
	return &defaultDetailFields
}

func DefaultDetailFieldsMap() *map[es.AnimeDetailField]bool {
	return &defaultDetailFieldsMap
}

func init() {
	defaultDetailFields = []es.AnimeDetailField{
		es.Id, es.AlternativeTitles, es.Title,
		es.Synopsis, es.Genres, es.Studios,
        es.Status, es.NumEpisodes, es.AverageEpisodeDuration,
        es.MainPicture,
	}

	defaultDetailFieldsMap = make(map[es.AnimeDetailField]bool)
	for _, v := range defaultDetailFields {
		defaultDetailFieldsMap[v] = true
	}

}
