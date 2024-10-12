package enums


/* const (
	Id = iota
	Title
	MainPicture
	AlternativeTitles
	StartDate
	EndDate
	Synopsis
	Mean
	Rank
	Popularity
	NumListUsers
	NumScoringUsers
	Nsfw
	CreatedAt
	UpdatedAt
	MediaType
	Status
	Genres
	MyListStatus
	NumEpisodes
	StartSeason
	Broadcast
	Source
	AverageEpisodeDuration
	Rating
	Pictures
	Background
	RelatedAnime
	RelatedManga
	Recommendations
	Studios
	Statistics
)
*/
// NOTE:
// stored in the order they are present in the constant's list
// therefore, iota value corresponds to the corresponding string value index.
var DetailsFields []string

func init() {
	DetailsFields = []string{
		"Id", "Title", "MainPicture", "AlternativeTitles", "StartDate",
		"EndDate", "Synopsis", "Mean", "Rank", "Popularity",
		"NumListUsers", "NumScoringUsers", "Nsfw", "CreatedAt", "UpdatedAt",
		"MediaType", "Status", "Genres", "MyListStatus", "NumEpisodes",
		"StartSeason", "Broadcast", "Source", "AverageEpisodeDuration", "Rating",
		"Pictures", "Background", "RelatedAnime", "RelatedManga", "Recommendations",
		"Studios", "Statistics",
    }
}
