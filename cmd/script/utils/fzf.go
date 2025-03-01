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

var imageDir = embedfiles.PreviewImageCacheDir
var dataDir = embedfiles.PreviewDataCacheDir

type FzfAnimeListParams struct {
	AnimeList *types.NativeAnimeList
	Limit     int
	Offset    *int
}

/*
shows the anime title's list using fzf

@return

	(animeId) int | -1 (if offset is changed)
	(error) fzf error
*/
func FzfAnimeList(p FzfAnimeListParams) (int, *types.AnimeListDataNode, error) {

	titleList := make([]string, 0)
	titleMap := make(map[string]struct {
		ID  int
		Idx int
	}, 0)

	for idx, val := range p.AnimeList.Data {
		cacheKey, fzfKey := GenerateAnimePreviewKeys(&val)

		titleMap[cacheKey] = struct {
			ID  int
			Idx int
		}{
			ID:  val.ID,
			Idx: idx,
		}

		titleList = append(titleList, strings.TrimSpace(fzfKey))

		if err := SaveAnimePreviewData(cacheKey, &val); err != nil {
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
	output, err := useFzf(titleList, "Search Results", previewScript, "Overview")
	if err != nil {
		return 0, &types.AnimeListDataNode{}, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
	}

	// find the output in titleList to get the index of the anime in the list so that we can use that node to get the cache key.

	selectedTitle := strings.TrimSpace(string(output))
	// Strip ANSI codes from selectedTitle to match the titleMap keys
	cleanTitle := StripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

	switch cleanTitle {
	case prevListStr:
		*(p.Offset) -= p.Limit
	case nextListStr:
		*(p.Offset) += p.Limit
	default:
		stk, found := titleMap[cleanTitle]
		if !found {
			fmt.Println("Selected title not found in the map.")
			panic("Title's shown in the fzf are not correctly mapped to their anime Id's")
		}
		return stk.ID, &p.AnimeList.Data[stk.Idx], nil
	}

	return -1, &types.AnimeListDataNode{}, nil
}

type FzfUserAnimeListParams struct {
	AnimeList *types.NativeUserAnimeList
	Limit     int
	Offset    *int
}

func FzfUserAnimeList(p FzfUserAnimeListParams) (int, *types.UserAnimeListDataNode, error) {

	titleList := make([]string, 0)
	titleMap := make(map[string]struct {
		ID  int
		Idx int
	}, 0)

	for idx, val := range p.AnimeList.Data {
		cacheKey, fzfKey := GenerateUserAnimePreviewKeys(&val.Node)

		titleMap[cacheKey] = struct {
			ID  int
			Idx int
		}{
			ID:  val.Node.ID,
			Idx: idx,
		}

		titleList = append(titleList, strings.TrimSpace(fzfKey))

		if err := SaveUserAnimePreviewData(cacheKey, &val); err != nil {
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
	output, err := useFzf(titleList, "User Anime List", previewScript, "Overview")
	if err != nil {
		return 0, &types.UserAnimeListDataNode{}, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
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
		stk, found := titleMap[cleanTitle]
		if !found {
			fmt.Println("Selected title not found in the map.")
			panic("Title's shown in the fzf are not correctly mapped to their anime Id's")
		}
		return stk.ID, &p.AnimeList.Data[stk.Idx], nil
	}
	return -1, &types.UserAnimeListDataNode{}, nil
}

type FzfRankingAnimeListParams struct {
	AnimeList   *types.NativeAnimeRanking
	Limit       int
	Offset      *int
	RankingType string
}

func FzfRankingAnimeList(p FzfRankingAnimeListParams) (int, *types.AnimeRankingDataNode, error) {

	titleList := make([]string, 0)
	titleMap := make(map[string]struct {
		ID  int
		Idx int
	}, 0)

	for idx, v := range p.AnimeList.Data {

		cacheKey, fzfKey := GenerateRankingAnimePreviewKeys(&v)

		titleMap[cacheKey] = struct {
			ID  int
			Idx int
		}{
			ID:  v.Node.ID,
			Idx: idx,
		}
		titleList = append(titleList, strings.TrimSpace(fzfKey))

		if err := SaveAnimePreviewData(cacheKey, &v.Node); err != nil {
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
	output, err := useFzf(titleList, fmt.Sprintf("Type: %s", p.RankingType), previewScript, "Overview")
	if err != nil {
		return 0, &types.AnimeRankingDataNode{}, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
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
		stk, found := titleMap[cleanTitle]
		if !found {
			fmt.Println("Selected title not found in the map.")
			panic("Title's shown in the fzf are not correctly mapped to their anime Id's")
		}
		return stk.ID, &p.AnimeList.Data[stk.Idx], nil
	}
	return -1, &types.AnimeRankingDataNode{}, nil
}

func FzfUserMenu(list []string, userD *types.NativeUserDetails) (enums.UserAnimeListStatus, error) {
	if err := SaveUserPreviewData(userD); err != nil {
		fmt.Println(err)
	}

	previewScript := GenerateUserPreviewScript()
	// Show user the list from which they can choose
	str, err := useFzf(list, "user info", previewScript, "User Info")
	if err != nil {
		return "", err
	}

	chosenListType, valid := enums.ParseUserAnimeListStatus(str)
	if !valid {
		return "", err
	}

	return chosenListType, nil
}

func useFzf(input []string, borderLabel string, previewScript string, previewLabel string) (string, error) {
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
		fmt.Sprintf("--preview-label=%s", previewLabel),
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
