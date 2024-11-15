package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/colors"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

type FzfAnimeListParams struct {
	AnimeList *types.NativeAnimeList
	Limit     int
	Offset    *int
}

var imageDir = os.Getenv("PREVIEW_IMAGE_CACHE_DIR")
var dataDir = os.Getenv("PREVIEW_DATA_CACHE_DIR")

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

		// download the image with this title and store it in the /tmp/mal-cli/images
		mainPicture := val.CustomFields["main_picture"]
		if mainPicture, ok := mainPicture.(map[string]interface{}); ok {
			medium, _ := mainPicture["medium"]
			url := fmt.Sprintf("%v", medium)
			fileName := strings.ReplaceAll(plainKeyStr.String(), " ", "")
			fileName = strings.ReplaceAll(fileName, "\t", "")
			fileName += filepath.Ext(url)

			go func() {
				// save preview images to cache
				if err := DownloadImage(url, imageDir+"/"+fileName); err != nil {
					fmt.Println(err)
					fmt.Println("error dowloading image")
				}
			}()
			go func() {
				// save preview data to cache
				fileName := strings.ReplaceAll(plainKeyStr.String(), " ", "")
				fileName = strings.ReplaceAll(fileName, "\t", "")
				SavePreviewData(dataDir+"/"+fileName, val)

			}()
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

	output, err := useFzf(titleList, "search results")
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

		mainPicture := val.CustomFields["main_picture"]
		if mainPicture, ok := mainPicture.(map[string]interface{}); ok {
			medium, _ := mainPicture["medium"]
			url := fmt.Sprintf("%v", medium)
			fileName := strings.ReplaceAll(plainKeyStr.String(), " ", "")
			fileName = strings.ReplaceAll(fileName, "\t", "")
			fileName += filepath.Ext(url)

			go func() {
				// save preview images to cache
				if err := DownloadImage(url, imageDir+"/"+fileName); err != nil {
					fmt.Println(err)
					fmt.Println("error dowloading image")
				}
			}()
			go func() {
				// save preview data to cache
				fileName := strings.ReplaceAll(plainKeyStr.String(), " ", "")
				fileName = strings.ReplaceAll(fileName, "\t", "")
				SavePreviewData(dataDir+"/"+fileName, val)

			}()
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

	output, err := useFzf(titleList, p.RankingType)
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

func useFzf(input []string, borderLabel string) (string, error) {

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
            title=$(echo {} | tr -d '[:space:]')
            show_image_previews="%s"
            if [ "${show_image_previews}" = "true" ];then
                if [ -s "%s/${title}.jpg" ]; then
                    fzf-preview "%s/${title}.jpg"
                elif [ -s "%s/${title}.png" ]; then
                    fzf-preview "%s/${title}.png"
                elif [ -s "%s/${title}.webp" ]; then
                    fzf-preview "%s/${title}.webp"
                else
                    echo Loading Image...
                fi
            fi
            if [ -s "%s/${title}" ]; then
                source "%s/${title}"
            else
                echo Loading Data...
            fi


            `,
			fzfPreview(),
			"true",
			imageDir, imageDir,
			imageDir, imageDir,
			imageDir, imageDir,
			dataDir, dataDir,
		),
	)
	fzf.Stdin = strings.NewReader(strings.Join(input, "\n"))

	output, err := fzf.Output()
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return string(output), nil
}
