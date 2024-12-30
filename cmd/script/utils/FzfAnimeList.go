package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	embedfiles "github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/embedFiles"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/colors"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

type FzfAnimeListParams struct {
	AnimeList *types.NativeAnimeList
	Limit     int
	Offset    *int
}

var imageDir = embedfiles.PreviewImageCacheDir
var dataDir = embedfiles.PreviewDataCacheDir

/*
shows the anime title's list using fzf

@return

	(animeId) int | -1 (if offset is changed)
	(error) fzf error
*/
func FzfAnimeList(p FzfAnimeListParams) (int, error) {

	titleList := make([]string, 0)
	titleMap := make(map[string]int, 0)
	for _, val := range p.AnimeList.Data {
		plainKeyStr := strings.Builder{}
		formattedKeyStr := strings.Builder{}

		// so to not mess with filepath
		val.Title = strings.ReplaceAll(val.Title, "\\", "-")
		val.Title = strings.ReplaceAll(val.Title, "/", "-")

		formattedKeyStr.WriteString(val.Title)
		plainKeyStr.WriteString(val.Title)

		startSeason := val.CustomFields["start_season"]
		if seasonData, ok := startSeason.(map[string]interface{}); ok {
			season, _ := seasonData["season"]
			year, _ := seasonData["year"]

			formattedKeyStr.WriteString("\t")
			plainKeyStr.WriteString(" ")
			plainKeyStr.WriteString(fmt.Sprintf("[ %v ~ %v ]", year, season))
			formattedKeyStr.WriteString(fmt.Sprintf("%s[ %v ~ %v ]%s", colors.Purple, year, season, colors.Reset))
		}

		titleMap[plainKeyStr.String()] = val.ID
		titleList = append(titleList, strings.TrimSpace(formattedKeyStr.String()))

		if err := SaveAnimePreviewData(plainKeyStr.String(), val); err != nil {
			fmt.Println(err)
		}
	}

	nextListStr := "Next List -->"
	prevListStr := "<-- Previous List"
	if *(p.Offset) > 0 {
		// previous list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, prevListStr, colors.Reset)))
	}
	if len(titleMap) != 0 {
		// next list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, nextListStr, colors.Reset)))
	}

	previewScript := GenerateAnimePreviewScript()
	output, err := useFzf(titleList, "search results", previewScript)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
	}

	selectedTitle := strings.TrimSpace(string(output))
	// Strip ANSI codes from selectedTitle to match the titleMap keys
	cleanTitle := StripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

	switch cleanTitle {
	case prevListStr:
		*(p.Offset) -= p.Limit
	case nextListStr:
		*(p.Offset) += p.Limit
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

type FzfUserAnimeListParams struct {
	AnimeList *types.NativeUserAnimeList
	Limit     int
	Offset    *int
}

func FzfUserAnimeList(p FzfUserAnimeListParams) (int, error) {

	titleList := make([]string, 0)
	titleMap := make(map[string]int, 0)
	for _, val := range p.AnimeList.Data {
		plainKeyStr := strings.Builder{}
		formattedKeyStr := strings.Builder{}

		// so to not mess with filepath
		val.Node.Title = strings.ReplaceAll(val.Node.Title, "\\", "-")
		val.Node.Title = strings.ReplaceAll(val.Node.Title, "/", "-")

		formattedKeyStr.WriteString(val.Node.Title)
		plainKeyStr.WriteString(val.Node.Title)

		startSeason := val.Node.CustomFields["start_season"]
		if seasonData, ok := startSeason.(map[string]interface{}); ok {
			season, _ := seasonData["season"]
			year, _ := seasonData["year"]

			formattedKeyStr.WriteString("\t")
			plainKeyStr.WriteString(" ")
			plainKeyStr.WriteString(fmt.Sprintf("[ %v ~ %v ]", year, season))
			formattedKeyStr.WriteString(fmt.Sprintf("%s[ %v ~ %v ]%s", colors.Purple, year, season, colors.Reset))
		}

		titleMap[plainKeyStr.String()] = val.Node.ID
		titleList = append(titleList, strings.TrimSpace(formattedKeyStr.String()))

		SaveUserAnimePreviewData(plainKeyStr.String(), val)
	}

	nextListStr := "Next List -->"
	prevListStr := "<-- Previous List"
	if *(p.Offset) > 0 {
		// previous list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, prevListStr, colors.Reset)))
	}
	if len(titleMap) != 0 {
		// next list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, nextListStr, colors.Reset)))
	}

	previewScript := GenerateAnimePreviewScript()
	output, err := useFzf(titleList, "search results", previewScript)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
	}

	selectedTitle := strings.TrimSpace(string(output))
	// Strip ANSI codes from selectedTitle to match the titleMap keys
	cleanTitle := StripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

	switch cleanTitle {
	case prevListStr:
		*(p.Offset) -= p.Limit
	case nextListStr:
		*(p.Offset) += p.Limit
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

type FzfRankingAnimeListParams struct {
	AnimeList   *types.NativeAnimeRanking
	Limit       int
	Offset      *int
	RankingType string
}

func FzfRankingAnimeList(p FzfRankingAnimeListParams) (int, error) {

	titleList := make([]string, 0)
	titleMap := make(map[string]int, 0)
	for _, v := range p.AnimeList.Data {
		val := v.Node

		plainKeyStr := strings.Builder{}
		formattedKeyStr := strings.Builder{}

		formattedKeyStr.WriteString(val.Title + "\t")
		plainKeyStr.WriteString(val.Title + " ")

		formattedKeyStr.WriteString(fmt.Sprintf("%sRank:%s %s%d%s", colors.Red, colors.Reset, colors.Blue, v.Ranking.Rank, colors.Reset))
		plainKeyStr.WriteString(fmt.Sprintf("Rank: %d", v.Ranking.Rank))

		titleMap[plainKeyStr.String()] = val.ID
		titleList = append(titleList, strings.TrimSpace(formattedKeyStr.String()))

		if err := SaveAnimePreviewData(plainKeyStr.String(), val); err != nil {
			fmt.Println(err)
		}
	}

	nextListStr := "Next List -->"
	prevListStr := "<-- Previous List"
	if *(p.Offset) > 0 {
		// previous list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, prevListStr, colors.Reset)))
	}
	if len(titleMap) != 0 {
		// next list
		titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, nextListStr, colors.Reset)))
	}

	previewScript := GenerateAnimePreviewScript()
	output, err := useFzf(titleList, p.RankingType, previewScript)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
	}

	selectedTitle := strings.TrimSpace(string(output))
	// Strip ANSI codes from selectedTitle to match the titleMap keys
	cleanTitle := StripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

	switch cleanTitle {
	case prevListStr:
		*(p.Offset) -= p.Limit
	case nextListStr:
		*(p.Offset) += p.Limit
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

func FzfUserMenu(list []string, userD *types.NativeUserDetails) (enums.UserAnimeListStatus, error) {
	if err := SaveUserPreviewData(userD); err != nil {
		fmt.Println(err)
	}

	previewScript := GenerateUserPreviewScript()
	// Show user the list from which they can choose
	str, err := useFzf(list, "user info", previewScript)
	if err != nil {
		return "", err
	}

	chosenListType, valid := enums.ParseUserAnimeListStatus(str)
	if !valid {
		return "", err
	}

	return chosenListType, nil
}

func useFzf(input []string, borderLabel string, previewScript string) (string, error) {
	fzf := exec.Command("fzf",
		"--no-sort",
		"--cycle",
		"--ansi",
		"+m",
		"--layout=reverse",
		"--border=rounded",
		"--preview-window=right:30%",
		"--wrap",
		fmt.Sprintf("--border-label=%s", borderLabel),
		fmt.Sprintf(`--preview=
                %s
                %s
            `,
			fzfPreview(),
			previewScript,
		),
	)
	fzf.Stdin = strings.NewReader(strings.Join(input, "\n"))

	output, err := fzf.Output()
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return string(output), nil
}
