package utils

import (
	"fmt"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/colors"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

// @return ("plainKey" for cache, "formattedKey" for fzf)
func GenerateAnimePreviewKeys(node *types.AnimeListDataNode) (string, string) {
	plainKeyStr := strings.Builder{}
	formattedKeyStr := strings.Builder{}

	// so to not mess with filepath
	node.Title = strings.ReplaceAll(node.Title, "\\", "-")
	node.Title = strings.ReplaceAll(node.Title, "/", "-")

	formattedKeyStr.WriteString(node.Title)
	plainKeyStr.WriteString(node.Title)

	startSeason := node.CustomFields["start_season"]
	if seasonData, ok := startSeason.(map[string]interface{}); ok {
		season, _ := seasonData["season"]
		year, _ := seasonData["year"]

		formattedKeyStr.WriteString("\t")
		plainKeyStr.WriteString(" ")
		plainKeyStr.WriteString(fmt.Sprintf("[ %v ~ %v ]", year, season))
		formattedKeyStr.WriteString(fmt.Sprintf("%s[ %v ~ %v ]%s", colors.Purple, year, season, colors.Reset))
	}

	return plainKeyStr.String(), formattedKeyStr.String()
}

// @return ("plainKey" for cache, "formattedKey" for fzf)
func GenerateUserAnimePreviewKeys(node *types.AnimeListDataNode) (string, string) {
	plainKeyStr := strings.Builder{}
	formattedKeyStr := strings.Builder{}

	// so to not mess with filepath
	node.Title = strings.ReplaceAll(node.Title, "\\", "-")
	node.Title = strings.ReplaceAll(node.Title, "/", "-")

	formattedKeyStr.WriteString(node.Title)
	plainKeyStr.WriteString(node.Title)

	startSeason := node.CustomFields["start_season"]
	if seasonData, ok := startSeason.(map[string]interface{}); ok {
		season, _ := seasonData["season"]
		year, _ := seasonData["year"]

		formattedKeyStr.WriteString("\t")
		plainKeyStr.WriteString(" ")
		plainKeyStr.WriteString(fmt.Sprintf("[ %v ~ %v ]", year, season))
		formattedKeyStr.WriteString(fmt.Sprintf("%s[ %v ~ %v ]%s", colors.Purple, year, season, colors.Reset))
	}

	return plainKeyStr.String(), formattedKeyStr.String()
}

// @return ("plainKey" for cache, "formattedKey" for fzf)
func GenerateRankingAnimePreviewKeys(node *types.AnimeRankingDataNode) (string, string) {
	plainKeyStr := strings.Builder{}
	formattedKeyStr := strings.Builder{}

	formattedKeyStr.WriteString(node.Node.Title + "\t")
	plainKeyStr.WriteString(node.Node.Title + " ")

	formattedKeyStr.WriteString(fmt.Sprintf("%sRank:%s %s%d%s", colors.Red, colors.Reset, colors.Blue, node.Ranking.Rank, colors.Reset))
	plainKeyStr.WriteString(fmt.Sprintf("Rank: %d", node.Ranking.Rank))

	return plainKeyStr.String(), formattedKeyStr.String()
}
