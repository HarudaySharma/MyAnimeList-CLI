package enums

type AnimeSeason string

const (
    Winter AnimeSeason = "winter"
    Spring AnimeSeason = "spring"
    Summer AnimeSeason = "summer"
    Fall AnimeSeason = "fall"
)

var animeseasons []AnimeSeason
var animeSeasonMap map[string]AnimeSeason

func AnimeSeasons() []AnimeSeason{
    return animeseasons
}

func ParseAnimeSeason(season string) (AnimeSeason, bool) {
    parsedSeason, found := animeSeasonMap[season]

    return parsedSeason, found
}


func init() {
    animeseasons = []AnimeSeason{
        Winter, Spring, Summer, Fall,
    }

    animeSeasonMap = make(map[string]AnimeSeason)
    for _, v := range(animeseasons) {
        animeSeasonMap[string(v)] = v
    }
}
