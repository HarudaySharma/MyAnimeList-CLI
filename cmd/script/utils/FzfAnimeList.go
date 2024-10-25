package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/colors"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

/*
shows the anime title's list using fzf

@return

	(animeId) int | -1 (if offset is changed)
	(error) fzf error
*/
func FzfAnimeList(animeList *types.NativeAnimeList, limit int, offset *int) (int, error) {

	titleList := make([]string, 0)
	titleMap := make(map[string]int, 0)
	for _, val := range animeList.Data {
		plainKeyStr := strings.Builder{}
		formattedKeyStr := strings.Builder{}

		formattedKeyStr.WriteString(val.Title + "\t")
		plainKeyStr.WriteString(val.Title + " ")

		startSeason := val.CustomFields["start_season"]
		if seasonData, ok := startSeason.(map[string]interface{}); ok {
			season, _ := seasonData["season"]
			year, _ := seasonData["year"]

			plainKeyStr.WriteString(fmt.Sprintf("[ %v ~ %v ]", year, season))
			formattedKeyStr.WriteString(fmt.Sprintf("%s[ %v ~ %v ]%s", colors.Purple, year, season, colors.Reset))
		}

		titleMap[plainKeyStr.String()] = val.ID
		titleList = append(titleList, strings.TrimSpace(formattedKeyStr.String()))
	}

	nextListStr := "Next List -->"
	prevListStr := "<-- Previous List"
	if *offset > 0 {
		// previous list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, prevListStr, colors.Reset)))
	}
	if len(titleMap) != 0 {
		// next list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, nextListStr, colors.Reset)))
	}

	output, err := useFzf(titleList)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
	}

	selectedTitle := strings.TrimSpace(string(output))
	// Strip ANSI codes from selectedTitle to match the titleMap keys
	cleanTitle := StripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

	switch cleanTitle {
	case prevListStr:
		*offset -= limit
	case nextListStr:
		*offset += limit
	default:
		animeId, found := titleMap[cleanTitle]
		if !found {
			fmt.Println("Selected title not found in the map.")
			panic("Title's shown in the fzf are not correctly mapped to their anime Id's")
		}
		return animeId, nil
	}
	return -1, nil
}


func useFzf(input []string) (string, error) {
	fzf := exec.Command("fzf", "--no-sort", "--cycle", "--ansi", "+m")
	fzf.Stdin = strings.NewReader(strings.Join(input, "\n"))

	output, err := fzf.Output()
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

    return string(output), nil
}
