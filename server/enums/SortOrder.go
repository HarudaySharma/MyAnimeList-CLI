package enums

import "strings"

type SortOption string

const (
	AnimeScore        SortOption = "anime_score"
	AnimeNumListUsers SortOption = "anime_num_list_users"
)

var sortOptions []SortOption
var sortOptionsMap map[string]SortOption

func SortOptions() []SortOption {
	return sortOptions
}

func ParseSortOptions(options []string) ([]SortOption, bool) {
	invalidOptionEncountered := false
	parsedOptions := make([]SortOption, 0)

	for _, option := range options {
        option = strings.Trim(option, " ")
        if option == "" {
            continue
        }

		opt, exists := sortOptionsMap[option]
		if !exists {
			invalidOptionEncountered = true
			continue
		}
		parsedOptions = append(parsedOptions, opt)
	}

	return parsedOptions, invalidOptionEncountered
}

func init() {
	sortOptions = []SortOption{
		AnimeNumListUsers,
		AnimeScore,
	}

	sortOptionsMap = make(map[string]SortOption)
	for _, option := range sortOptions {
		sortOptionsMap[string(option)] = option
	}
}
