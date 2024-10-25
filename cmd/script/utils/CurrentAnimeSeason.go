package utils

import (
	"time"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
)

func CurrentAnimeSeason() enums.AnimeSeason {
	month := time.Now().Month()

	switch month {
	case time.January, time.February, time.March:
		return enums.Winter
	case time.April, time.May, time.June:
		return enums.Spring
	case time.July, time.August, time.September:
		return enums.Summer
	case time.October, time.November, time.December:
		return enums.Fall
	}

	panic("Something wrong with CurrentAnimeSeason()")
}
