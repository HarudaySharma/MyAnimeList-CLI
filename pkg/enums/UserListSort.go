package enums

import "strings"

type UserAnimeListSortOption string

const (
	UALSort_ListScore      UserAnimeListSortOption = "list_score"
	UALSort_ListUpdatedAt  UserAnimeListSortOption = "list_updated_at"
	UALSort_AnimeTitle     UserAnimeListSortOption = "anime_title"
	UALSort_AnimeStartDate UserAnimeListSortOption = "anime_start_date"
)

var userAnimeListSortOptions []UserAnimeListSortOption
var userAnimeListSortOptionMap map[string]UserAnimeListSortOption

func UserAnimeListSortOptions() []UserAnimeListSortOption {
	return userAnimeListSortOptions
}

func ParseUserAnimeListSortOptions(options []string) ([]UserAnimeListSortOption, bool) {
	invalidOptionEncountered := false
	parsedOptions := make([]UserAnimeListSortOption, 0)

	for _, option := range options {
		option = strings.Trim(option, " ")
		if option == "" {
			continue
		}

		opt, exists := userAnimeListSortOptionMap[option]
		if !exists {
			invalidOptionEncountered = true
			continue
		}
		parsedOptions = append(parsedOptions, opt)
	}

	return parsedOptions, invalidOptionEncountered
}

func init() {
	userAnimeListSortOptions = []UserAnimeListSortOption{
		UALSort_AnimeTitle,
		UALSort_AnimeStartDate,
		UALSort_ListScore,
		UALSort_ListUpdatedAt,
	}

	userAnimeListSortOptionMap = make(map[string]UserAnimeListSortOption)
	for _, option := range userAnimeListSortOptions {
		userAnimeListSortOptionMap[string(option)] = option
	}
}
