package enums

type Ranking string

const (
	RankingAll          Ranking = "all"
	RankingAiring       Ranking = "airing"
	RankingUpcoming     Ranking = "upcoming"
	RankingTV           Ranking = "tv"
	RankingOVA          Ranking = "ova"
	RankingMovie        Ranking = "movie"
	RankingSpecial      Ranking = "special"
	RankingByPopularity Ranking = "popularity"
	RankingByFavorite   Ranking = "favorite"
)

var ranking []Ranking

func AnimeRanking() []Ranking {
	return ranking
}

func ParseAnimeRaking(rankStr string) (Ranking, bool) {
	invalid := true
	var rank Ranking
	for _, v := range ranking {
		if string(v) == rankStr {
			rank = v
			invalid = false
			break
		}
	}

	return rank, invalid
}

func init() {
	ranking = []Ranking{
		RankingAll, RankingAiring, RankingUpcoming, RankingTV, RankingOVA,
		RankingMovie, RankingSpecial, RankingByPopularity, RankingByFavorite,
	}
}
