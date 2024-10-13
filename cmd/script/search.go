package script

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/ui"
	u "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
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
				fmt.Print("Enter the anime title [atleast 3 letters word]: ")
				input, _ := reader.ReadString('\n')
				query = strings.TrimSpace(input)
			}
		} else {
			query = strings.Join(args, " ")
			query = strings.TrimSpace(query)
		}

		encodedQuery := url.QueryEscape(query)
		var animeList types.NativeAnimeList
		err := u.GetAnimeList(&animeList, encodedQuery, 100, 0, []es.AnimeDetailField{
			es.StartSeason,
		})
		if err != nil {
			fmt.Println(err)
			return
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
				formattedKeyStr.WriteString(fmt.Sprintf("\033[35m[ %v ~ %v ]\033[0m", year, season))
			}

			titleMap[plainKeyStr.String()] = val.ID
			titleList = append(titleList, strings.TrimSpace(formattedKeyStr.String()))
		}

		output, err := u.UseFzf(titleList)
		if err != nil {
			fmt.Printf("error using fzf \n %v\n", err)
			return
		}

		selectedTitle := strings.TrimSpace(string(output))
		// Strip ANSI codes from selectedTitle to match the titleMap keys
		cleanTitle := u.StripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

		animeId, found := titleMap[cleanTitle]
		if !found {
			fmt.Println("Selected title not found in the map.")
			return
		}

		detailsIdxs, _ := cmd.Flags().GetIntSlice("d")

		detailFields := make([]es.AnimeDetailField, 0)
		detailFields = append(detailFields, *enums.DefaultDetailFields()...)
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

func init() {

	availableOptions := u.ConvertToCommaSeperatedString(u.ConvertToCommaSeperatedStringParams[es.AnimeDetailField]{
		Data:            es.EveryDetailField(),
		MaxLineLen:      38,
		SpaceAfterComma: true,
	})
	totalOptions := len(es.EveryDetailField())

	searchCmd.PersistentFlags().IntSlice("d", []int{}, strings.TrimSpace(fmt.Sprintf(`
        Specify which anime detail you want

            Available Options:
                %s
            Note:
                options value are from 0..%d
                    0 => %s
                    .......
                    %d => %s

            Usage:
                ani-cli search "anime title" --d=1,2,31
                ani-cli search "anime title" --d 1,2,31
        `,
		availableOptions,
		totalOptions-1,
		es.EveryDetailField()[0],
		totalOptions-1,
		es.EveryDetailField()[totalOptions-1],
	)))
	rootCmd.AddCommand(searchCmd)
}
