package enums

type Ranking string

const (
	All          Ranking = "all"
	Airing       Ranking = "airing"
	Upcoming     Ranking = "upcoming"
	TV           Ranking = "tv"
	OVA          Ranking = "ova"
	Movie        Ranking = "movie"
	Special      Ranking = "special"
	ByPopularity Ranking = "bypopularity"
	Favorite     Ranking = "favorite"
)

var ranking []Ranking

func AnimeRanking() []Ranking {
	return ranking
}

func ParseAnimeRaking(raking string) (Ranking, bool) {
    matched := false
    var rank Ranking
    for _, v := range(ranking) {
        if string(v) == raking {
            rank = v
            matched = true
            break
        }
    }

    return rank, matched
}

func init() {
	ranking = []Ranking{
		All, Airing, Upcoming, TV, OVA,
		Movie, Special, ByPopularity, Favorite,
	}
}
