package types

import (
	"time"

	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
)

type AlternativeTitles struct {
	EN       string   `json:"en"`
	JA       string   `json:"ja"`
	Synonyms []string `json:"synonyms"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Picture struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type MALDataNode struct {
	ID           int                    `json:"id"`
	MainPicture  Picture                `json:"main_picture"`
	Title        string                 `json:"title"`
	CustomFields map[string]interface{} `json:"-"`
}

type Recommendation struct {
	Node               MALDataNode `json:"node"`
	NumRecommendations int         `json:"num_recommendations"`
}

type RelatedAnimeNode struct {
	ID          int     `json:"id"`
	MainPicture Picture `json:"main_picture"`
	Title       string  `json:"title"`
}

type RelatedAnime struct {
	Node                  RelatedAnimeNode `json:"node"`
	RelationType          string           `json:"relation_type"`
	RelationTypeFormatted string           `json:"relation_type_formatted"`
}

type Broadcast struct {
	DayOfTheWeek string `json:"day_of_the_week"` // in Japan time
	StartTime    string `json:"start_time"`      // 24 hrs format (not listed in docs)
}

type Statistics struct {
	NumListUsers int64 `json:"num_list_users"`
	Status       struct {
		Completed   interface{} `json:"completed"`     // can be string or int
		Dropped     interface{} `json:"dropped"`       // can be string or int
		OnHold      interface{} `json:"on_hold"`       // can be string or int
		PlanToWatch interface{} `json:"plan_to_watch"` // can be string or int
		Watching    interface{} `json:"watching"`      // can be string or int
	} `json:"status"`
}

type Studio struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Page struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type MALAnimeList struct {
	Data []struct {
		Node MALDataNode `json:"node"`
	} `json:"data"`
	/* Paging struct {
		Next string `json:"next"`
	} `json:"paging"` */
	Paging Page `json:"paging"`
}

type MALAnimeDetails struct {
	AlternativeTitles      AlternativeTitles `json:"alternative_titles"`
	AverageEpisodeDuration int64             `json:"average_episode_duration"`
	Background             string            `json:"background"`
	Broadcast              Broadcast         `json:"broadcast"`
	CreatedAt              time.Time         `json:"created_at"`
	EndDate                string            `json:"end_date"`
	Genres                 []Genre           `json:"genres"`
	ID                     int               `json:"id"`
	MainPicture            Picture           `json:"main_picture"`
	Mean                   float64           `json:"mean"`
	MediaType              string            `json:"media_type"`
	NSFW                   string            `json:"nsfw"`
	NumEpisodes            int               `json:"num_episodes"`
	NumListUsers           int64             `json:"num_list_users"`
	NumScoringUsers        int64             `json:"num_scoring_users"`
	Pictures               []Picture         `json:"pictures"`
	Popularity             int               `json:"popularity"`
	Rank                   int               `json:"rank"`
	Rating                 string            `json:"rating"`
	Recommendations        []Recommendation  `json:"recommendations"`
	RelatedAnime           []RelatedAnime    `json:"related_anime"`
	Source                 string            `json:"source"`
	StartDate              string            `json:"start_date"`
	StartSeason            struct {
		Season string `json:"season"`
		Year   int    `json:"year"`
	} `json:"start_season"`
	Statistics Statistics `json:"statistics"`
	Status     string     `json:"status"`
	Studios    []Studio   `json:"studios"`
	Synopsis   string     `json:"synopsis"`
	Title      string     `json:"title"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type MALAnimeRanking struct {
	Data []struct {
		Node    MALDataNode `json:"node"`
		Ranking struct {
			Rank int `json:"rank"`
		} `json:"ranking"`
	} `json:"data"`
	/* Paging struct {
		Next string `json:"next"`
	} `json:"paging"` */
	Paging Page `json:"paging"`
}

type MALSeasonalAnime struct {
	Data []struct {
		Node MALDataNode `json:"node"`
	} `json:"data"`
	Paging Page `json:"paging"`
}

/**********************************************/
/**********************************************/
/**************** NATIVE TYPES ****************/
/**********************************************/
/**********************************************/

type AnimeListDataNode struct {
	ID           int                    `json:"id"`
	Title        string                 `json:"title"`
	CustomFields map[string]interface{} `json:""`
}

type AnimeRankingDataNode struct {
	Node    AnimeListDataNode
	Ranking struct {
		Rank int `json:"rank"`
	} `json:"ranking"`
}

type NativeRecommendation struct {
	Node               AnimeListDataNode `json:"node"`
	NumRecommendations int               `json:"num_recommendations"`
}

type NativeAnimeList struct {
	Data []AnimeListDataNode `json:"data"`
	/* Paging struct {
		Next string `json:"next"`
	} `json:"paging"` */
	Paging Page `json:"paging"`
}

type NativeRelatedAnime struct {
	Node                  AnimeListDataNode `json:"node"`
	RelationType          string            `json:"relation_type"`
	RelationTypeFormatted string            `json:"relation_type_formatted"`
}

type NativeAnimeDetails struct {
	AlternativeTitles      AlternativeTitles      `json:"alternative_titles"`
	AverageEpisodeDuration int64                  `json:"average_episode_duration"` // in seconds
	Background             string                 `json:"background"`
	Broadcast              Broadcast              `json:"broadcast"`
	CreatedAt              time.Time              `json:"created_at"`
	EndDate                string                 `json:"end_date"`
	Genres                 []Genre                `json:"genres"`
	ID                     int                    `json:"id"`
	MainPicture            Picture                `json:"main_picture"`
	Mean                   float64                `json:"mean"`
	MediaType              string                 `json:"media_type"`
	NSFW                   string                 `json:"nsfw"`
	NumEpisodes            int                    `json:"num_episodes"`
	NumListUsers           int64                  `json:"num_list_users"`
	NumScoringUsers        int64                  `json:"num_scoring_users"`
	Pictures               []Picture              `json:"pictures"`
	Popularity             int                    `json:"popularity"`
	Rank                   int                    `json:"rank"`
	Rating                 string                 `json:"rating"`
	Recommendations        []NativeRecommendation `json:"recommendations"`
	RelatedAnime           []NativeRelatedAnime   `json:"related_anime"`
	Source                 string                 `json:"source"`
	StartDate              string                 `json:"start_date"`
	StartSeason            struct {
		Season string `json:"season"`
		Year   int    `json:"year"`
	} `json:"start_season"`
	Statistics Statistics `json:"statistics"`
	Status     string     `json:"status"`
	Studios    []Studio   `json:"studios"`
	Synopsis   string     `json:"synopsis"`
	Title      string     `json:"title"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type NativeAnimeDetails_Basic struct {
	AlternativeTitles      AlternativeTitles `json:"alternative_titles"`
	AverageEpisodeDuration int               `json:"average_episode_duration"`
	Broadcast              Broadcast         `json:"broadcast"`
	CreatedAt              time.Time         `json:"created_at"`
	EndDate                string            `json:"end_date"`
	Genres                 []Genre           `json:"genres"`
	ID                     int               `json:"id"`
	MainPicture            Picture           `json:"main_picture"`
	MediaType              string            `json:"media_type"`
	NumEpisodes            int               `json:"num_episodes"`
	Rank                   int               `json:"rank"`
	Rating                 string            `json:"rating"`
	Source                 string            `json:"source"`
	StartSeason            struct {
		Season string `json:"season"`
		Year   int    `json:"year"`
	} `json:"start_season"`
	Status   string `json:"status"`
	Synopsis string `json:"synopsis"`
	Title    string `json:"title"`
}

type NativeAnimeDetails_Advanced struct {
	AlternativeTitles      AlternativeTitles      `json:"alternative_titles"`
	AverageEpisodeDuration int                    `json:"average_episode_duration"`
	Background             string                 `json:"background"`
	Broadcast              Broadcast              `json:"broadcast"`
	CreatedAt              time.Time              `json:"created_at"`
	EndDate                string                 `json:"end_date"`
	Genres                 []Genre                `json:"genres"`
	ID                     int                    `json:"id"`
	MainPicture            Picture                `json:"main_picture"`
	MediaType              string                 `json:"media_type"`
	NumEpisodes            int                    `json:"num_episodes"`
	NumListUsers           int                    `json:"num_list_users"`
	NumScoringUsers        int                    `json:"num_scoring_users"`
	Rank                   int                    `json:"rank"`
	Rating                 string                 `json:"rating"`
	Recommendations        []NativeRecommendation `json:"recommendations"`
	RelatedAnime           []NativeRelatedAnime   `json:"related_anime"`
	Source                 string                 `json:"source"`
	StartDate              string                 `json:"start_date"`
	StartSeason            struct {
		Season string `json:"season"`
		Year   int    `json:"year"`
	} `json:"start_season"`
	Statistics Statistics `json:"statistics"`
	Status     string     `json:"status"`
	Studios    []Studio   `json:"studios"`
	Synopsis   string     `json:"synopsis"`
	Title      string     `json:"title"`
}

type NativeAnimeRanking struct {
	Data []AnimeRankingDataNode `json:"data"`
	/* Paging struct {
		Next string `json:"next"`
	} */
	Paging Page `json:"paging"`
}

type NativeSeasonalAnime struct {
	Data   []AnimeListDataNode `json:"data"`
	Paging Page                `json:"paging"`
}

// ***GROUPING RELATED TYPES***
type AnimeDetails interface {
	NativeAnimeDetails | NativeAnimeDetails_Basic | NativeAnimeDetails_Advanced
}

// TODO: make use of this
type AnimeList interface {
	NativeSeasonalAnime | NativeAnimeList | NativeAnimeRanking
}

/*********************************/
/*********************************/
/****** USER SPECIFIC TYPES ******/
/*********************************/
/*********************************/

type MALUserDetails struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	Picture         string    `json:"picture"`
	Location        string    `json:"location"`
	JoinedAt        time.Time `json:"joined_at"`
	AnimeStatistics struct {
		NumItemsWatching    float64 `json:"num_items_watching"`
		NumItemsCompleted   float64 `json:"num_items_completed"`
		NumItemsOnHold      float64 `json:"num_items_on_hold"`
		NumItemsDropped     float64 `json:"num_items_dropped"`
		NumItemsPlanToWatch float64 `json:"num_items_plan_to_watch"`
		NumItems            float64 `json:"num_items"`
		NumDaysWatched      float64 `json:"num_days_watched"`
		NumDaysWatching     float64 `json:"num_days_watching"`
		NumDaysCompleted    float64 `json:"num_days_completed"`
		NumDaysOnHold       float64 `json:"num_days_on_hold"`
		NumDaysDropped      float64 `json:"num_days_dropped"`
		NumDays             float64 `json:"num_days"`
		NumEpisodes         float64 `json:"num_episodes"`
		NumTimesRewatched   float64 `json:"num_times_rewatched"`
		MeanScore           float64 `json:"mean_score"  `
	} `json:"anime_statistics"`

	Error string `json:"error"`
}

type NativeUserDetails struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	Picture         string    `json:"picture"`
	Location        string    `json:"location"`
	JoinedAt        time.Time `json:"joined_at"`
	AnimeStatistics struct {
		NumItemsWatching    float64 `json:"num_items_watching"`
		NumItemsCompleted   float64 `json:"num_items_completed"`
		NumItemsOnHold      float64 `json:"num_items_on_hold"`
		NumItemsDropped     float64 `json:"num_items_dropped"`
		NumItemsPlanToWatch float64 `json:"num_items_plan_to_watch"`
		NumItems            float64 `json:"num_items"`
		NumDaysWatched      float64 `json:"num_days_watched"`
		NumDaysWatching     float64 `json:"num_days_watching"`
		NumDaysCompleted    float64 `json:"num_days_completed"`
		NumDaysOnHold       float64 `json:"num_days_on_hold"`
		NumDaysDropped      float64 `json:"num_days_dropped"`
		NumDays             float64 `json:"num_days"`
		NumEpisodes         float64 `json:"num_episodes"`
		NumTimesRewatched   float64 `json:"num_times_rewatched"`
		MeanScore           float64 `json:"mean_score"  `
	} `json:"anime_statistics"`
}

type MALUserAnimeList struct {
	Data []struct {
		Node        MALDataNode        `json:"node"`
		AnimeStatus MALUserAnimeStatus `json:"list_status"`
	} `json:"data"`

	Paging Page `json:"paging"`
}

type MALUserAnimeStatus struct {
	Status               enums.UserAnimeListStatus `json:"status"`
	Score                int8                      `json:"score"` // 0-10
	NumEpisodesWatched   int16                     `json:"num_episodes_watched"`// NOTE: in docs the json field is "num_watched_episodes" which is wrong
	IsRewatching         bool                      `json:"is_rewatching"`
	UpdatedAt            time.Time                 `json:"updated_at"`
	Priority             int                       `json:"priority"` // 0-2
	NumberTimesRewatched int                       `json:"num_times_rewatched"`
	RewatchValue         int                       `json:"rewatch_value"` // 0-5
	Tags                 []string                  `json:"tags"`
	Comments             string                    `json:"comments"`
}

type NativeUserAnimeStatus struct {
	Status             enums.UserAnimeListStatus `json:"status"`
	Score              int8                      `json:"score"` // 0-10
    NumWatchedEpisodes int16                     `json:"num_watched_episodes"` // NOTE: not named identical to the source json field as when updating the status they require you to send a "num_watched_episodes" field instead of "num_episodes_watched"
	IsRewatching       bool                      `json:"is_rewatching"`
	UpdatedAt          time.Time                 `json:"updated_at"`
}

type UserAnimeListDataNode struct {
	Node        AnimeListDataNode     `json:"node"`
	AnimeStatus NativeUserAnimeStatus `json:"list_status"`
}

type NativeUserAnimeList struct {
	Data   []UserAnimeListDataNode `json:"data"`
	Paging Page                    `json:"paging"`
}
