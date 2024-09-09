package utils


type AnimeDetailField string

const (
	Id                     AnimeDetailField = "id"
	MainPicture            AnimeDetailField = "main_picture"
	AlternativeTitles      AnimeDetailField = "alternative_titles"
	StartDate              AnimeDetailField = "start_date"
	EndDate                AnimeDetailField = "end_date"
	Synopsis               AnimeDetailField = "synopsis"
	Mean                   AnimeDetailField = "mean"
	Rank                   AnimeDetailField = "rank"
	Popularity             AnimeDetailField = "popularity"
	NumListUsers           AnimeDetailField = "num_list_users"
	NumScoringUsers        AnimeDetailField = "num_scoring_users"
	Nsfw                   AnimeDetailField = "nsfw"
	CreatedAt              AnimeDetailField = "created_at"
	UpdatedAt              AnimeDetailField = "updated_at"
	MediaType              AnimeDetailField = "media_type"
	Status                 AnimeDetailField = "status"
	Genres                 AnimeDetailField = "genres"
	MyListStatus           AnimeDetailField = "my_list_status"
	NumEpisodes            AnimeDetailField = "num_episodes"
	StartSeason            AnimeDetailField = "start_season"
	Broadcast              AnimeDetailField = "broadcast"
	Source                 AnimeDetailField = "source"
	AverageEpisodeDuration AnimeDetailField = "average_episode_duration"
	Rating                 AnimeDetailField = "rating"
	Pictures               AnimeDetailField = "pictures"
	Background             AnimeDetailField = "background"
	RelatedAnime           AnimeDetailField = "related_anime"
	RelatedManga           AnimeDetailField = "related_manga"
	Recommendations        AnimeDetailField = "recommendations"
	Studios                AnimeDetailField = "studios"
	Statistics             AnimeDetailField = "statistics"
)

var basicDetailFields []AnimeDetailField
var allFields []AnimeDetailField
var allFieldsMap map[string]AnimeDetailField

func BasicDetailFields() []AnimeDetailField {
	return basicDetailFields
}

func AllFields() []AnimeDetailField {
	return allFields
}

func ParseDetailsField(fields []string) ([]AnimeDetailField, bool) {
	parsedFields := make([]AnimeDetailField, 0)
    invalidFieldEncountered := false
	for _, field := range fields {
		f, exists := allFieldsMap[field]
		if !exists {
            invalidFieldEncountered = true
			continue
			// NOT SURE
			// should just ignore the fields that are invalid 
			// return []AnimeDetailField{}, errors.New(fmt.Sprintf("invalid field {%s}", field))
		}
		parsedFields = append(parsedFields, f)
	}

	return parsedFields, invalidFieldEncountered
}

func init() {
	basicDetailFields = []AnimeDetailField{
		Id, AlternativeTitles, StartDate, EndDate, Synopsis, Rank,
		Popularity, NumListUsers, NumScoringUsers, Nsfw, CreatedAt, UpdatedAt,
		Status, Genres, NumEpisodes, StartSeason, Source,
		AverageEpisodeDuration, Rating, RelatedAnime,
		Studios, Statistics,
	}

	allFields = []AnimeDetailField{
		Id, MainPicture, AlternativeTitles, StartDate, EndDate, Synopsis, Mean, Rank,
		Popularity, NumListUsers, NumScoringUsers, Nsfw, CreatedAt, UpdatedAt, MediaType,
		Status, Genres, MyListStatus, NumEpisodes, StartSeason, Broadcast, Source,
		AverageEpisodeDuration, Rating, Pictures, Background, RelatedAnime, RelatedManga,
		Recommendations, Studios, Statistics,
	}

	allFieldsMap = make(map[string]AnimeDetailField)

	for _, f := range allFields {
		allFieldsMap[string(f)] = f
	}
}
