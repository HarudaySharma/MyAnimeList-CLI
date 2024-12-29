package enums

import (
	"strings"
)

type UserAnimeListStatus string

const (
	ULS_ALL    UserAnimeListStatus = "all"
	ULS_Watching    UserAnimeListStatus = "watching"
	ULS_Completed   UserAnimeListStatus = "completed"
	ULS_OnHold      UserAnimeListStatus = "on_hold"
	ULS_Dropped     UserAnimeListStatus = "dropped"
	ULS_PlanToWatch UserAnimeListStatus = "plan_to_watch"
)

var userAnimeListStatuses []UserAnimeListStatus
var userAnimeListStatusMap map[string]UserAnimeListStatus

func UserAnimeListStatuses() []UserAnimeListStatus {
	return userAnimeListStatuses
}

func ParseUserAnimeListStatus(option string) (UserAnimeListStatus, bool) {
	option = strings.TrimSpace(option)
	if option == "" {
		return "", false
	}

    opt, exists := userAnimeListStatusMap[option]

	return opt, exists

}

func init() {
	userAnimeListStatuses = []UserAnimeListStatus{
        ULS_ALL,
        ULS_Watching,
		ULS_PlanToWatch,
		ULS_OnHold,
		ULS_Dropped,
		ULS_Completed,
	}

	userAnimeListStatusMap = make(map[string]UserAnimeListStatus)
	for _, option := range userAnimeListStatuses {
		userAnimeListStatusMap[string(option)] = option
	}
}
