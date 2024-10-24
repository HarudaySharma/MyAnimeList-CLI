package script

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	 e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/ui"
	u "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/colors"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search anime by title",
	Run: func(cmd *cobra.Command, args []string) {
		var query string
		if len(args) == 0 || len(args[0]) < 3 {
			reader := bufio.NewReader(os.Stdin)
			for len(query) < 3 {
				fmt.Print(
					colors.Blue, "Enter the anime title", colors.Reset,
					colors.Red, " [atleast 3 letters word]: ", colors.Reset,
				)
				input, _ := reader.ReadString('\n')
				query = strings.TrimSpace(input)
			}
		} else {
			query = strings.Join(args, " ")
			query = strings.TrimSpace(query)
		}

		offset := 0
		limit, err := cmd.Flags().GetInt("l")
		if err != nil {
			limit = 10 // don't panic
		}

		var animeId int
		for {
			animeId, err = showAnimeList(&query, &limit, &offset)
			if err != nil {
				if strings.Contains(err.Error(), "130") { // 130 for ESC in FZF
					os.Exit(0)
				}

				fmt.Printf("%v\n****Unexpected Error, Please try again!!****", err)
				fmt.Println()

				reader := bufio.NewReader(os.Stdin)
				query = ""
				for len(query) < 3 {
					fmt.Print(
						colors.Blue, "Enter the anime title", colors.Reset,
						colors.Red, " [atleast 3 letters word]: ", colors.Reset,
					)
					input, _ := reader.ReadString('\n')
					query = strings.TrimSpace(input)
				}
				continue
			}
			break
		}

		detailsIdxs, _ := cmd.Flags().GetIntSlice("d")

		detailFields := make([]es.AnimeDetailField, 0)
		detailFields = append(detailFields, *e.DefaultDetailFields()...)
		detailFields = append(detailFields, u.MapIndicesToDetailFields(detailsIdxs)...)

		var animeDetails types.NativeAnimeDetails
		u.GetAnimeDetails(&animeDetails, animeId, "custom", detailFields)

		animeDetailsUI := ui.AnimeDetailsUI{
			Details:      &animeDetails,
			DetailFields: &detailFields,
		}

		if err := tview.NewApplication().SetRoot(animeDetailsUI.CreateLayout(), true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

/*
@return

	animeId: int
	error: fzf error
*/
func showAnimeList(query *string, limit, offset *int) (int, error) {
	var animeList types.NativeAnimeList
	for {
		err := u.GetAnimeList(&animeList, *query, *limit, *offset, []es.AnimeDetailField{
			es.StartSeason,
		})
		if err != nil {
			fmt.Println(err)
			return 0, errors.New(fmt.Sprintf("%v\n****ERROR GETTING ANIMELIST**** ", err))
		}

		// show the list with fzf
		// TODO: give an option for [next list & previous list]

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
		// next list
		if len(titleMap) != 0 {
			titleList = append(titleList, strings.TrimSpace(fmt.Sprintf("%s%s%s", colors.Purple, nextListStr, colors.Reset)))
		}

		output, err := u.UseFzf(titleList)
		if err != nil {
			return 0, errors.New(fmt.Sprintf("error using fzf \n %v\n", err))
		}

		selectedTitle := strings.TrimSpace(string(output))
		// Strip ANSI codes from selectedTitle to match the titleMap keys
		cleanTitle := u.StripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

		switch cleanTitle {
		case prevListStr:
			*offset -= *limit
		case nextListStr:
			*offset += *limit
		default:
			animeId, found := titleMap[cleanTitle]
			if !found {
				fmt.Println("Selected title not found in the map.")
				continue
			}
			return animeId, nil
		}
	}
}

func init() {

	searchCmd.PersistentFlags().Int("l", 10, strings.TrimSpace(fmt.Sprintf(`
        Specify length of anime list { 1 - 100 }
        `,
	)))

	availableOptionsStr := strings.Builder{}
	availableOptionsStr.WriteString("\n\t\t")
	for i, option := range es.EveryDetailField() {
		availableOptionsStr.WriteString(fmt.Sprintf("%d => %s", i, option))
		availableOptionsStr.WriteString("\n\t\t")
	}

	searchCmd.PersistentFlags().IntSlice("d", []int{}, strings.TrimSpace(fmt.Sprintf(`
        Specify which anime detail you want

            Available Options: %s
            Usage:
                ani-cli search "anime title" --d=1,2,31
                ani-cli search "anime title" --d 1,2,31
        `,
		availableOptionsStr.String(),
	)))
	rootCmd.AddCommand(searchCmd)

}
