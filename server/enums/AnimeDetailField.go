package enums

type AnimeDetailField string

const (
	Id                     AnimeDetailField = "id"
    Title                  AnimeDetailField = "title"
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
var advancedDetailFields []AnimeDetailField
var everyDetailFields []AnimeDetailField
var allFieldsMap map[string]AnimeDetailField

func BasicDetailFields() []AnimeDetailField {
	return basicDetailFields
}

func AdvancedDetailFields() []AnimeDetailField {
	return advancedDetailFields
}

func EveryDetailField() []AnimeDetailField {
	return everyDetailFields
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
		Id, AlternativeTitles, StartDate, Synopsis, Rank,
		CreatedAt, Status, Genres, NumEpisodes, StartSeason,
		AverageEpisodeDuration, Rating,
	}

	advancedDetailFields = []AnimeDetailField{
		Id, AlternativeTitles, StartDate, Synopsis, Rank,
		NumListUsers, NumScoringUsers, CreatedAt,
		Status, Genres, NumEpisodes, StartSeason, Source,
		AverageEpisodeDuration, Rating, RelatedAnime,
		Studios, Statistics,
	}

	everyDetailFields = []AnimeDetailField{
		Id, Title, MainPicture, AlternativeTitles, StartDate, EndDate, Synopsis, Mean, Rank,
		Popularity, NumListUsers, NumScoringUsers, Nsfw, CreatedAt, UpdatedAt, MediaType,
		Status, Genres, MyListStatus, NumEpisodes, StartSeason, Broadcast, Source,
		AverageEpisodeDuration, Rating, Pictures, Background, RelatedAnime, RelatedManga,
		Recommendations, Studios, Statistics,
	}

	allFieldsMap = make(map[string]AnimeDetailField)

	for _, f := range everyDetailFields {
		allFieldsMap[string(f)] = f
	}
}
