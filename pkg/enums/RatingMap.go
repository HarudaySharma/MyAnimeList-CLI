package enums

var ratingMap map[string]string

func RatingMap() map[string]string {
	return ratingMap
}

func init() {
	ratings := []string{
		"g", "pg",
		"pg_13", "r",
		"rx", "r+",
	}

	descs := []string{
		"G - All Ages",
		"PG - Children",
		"PG-13 - Teens 13 and Older",
		"R - 17+ (violence & profanity)",
		"Rx - Hentai",
		"R+ - Profanity & Mild Nudity",
	}

	ratingMap = make(map[string]string)

	for i, rating := range ratings {
		ratingMap[rating] = descs[i]
	}

}
