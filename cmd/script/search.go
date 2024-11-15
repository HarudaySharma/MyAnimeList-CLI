package script

import (
	"bufio"
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
		if len(args) != 0 && len(args[0]) >= 3 {
			query = strings.Join(args, " ")
			query = strings.TrimSpace(query)
		}

		offset := 0
		limit, err := rootCmd.Flags().GetInt("list-size")
		if err != nil {
			limit = 10 // don't panic
		}

		if limit > es.MAX_SEARCH_LIST_SIZE {
			limit = es.MAX_SEARCH_LIST_SIZE
		}

		var animeList types.NativeAnimeList
		var animeId int = -1 // will send a request to server every time it is "-1"

		detailFields := make([]es.AnimeDetailField, 0)
		detailFields = append(detailFields, *e.PreviewDetailFields()...)
		detailFields = append(detailFields, es.Title, es.StartSeason)

		for {
			for {
				for len(query) < 3 {
					reader := bufio.NewReader(os.Stdin)
					fmt.Print(
						colors.Blue, "Enter the anime title", colors.Reset,
						colors.Red, " [atleast 3 letters word]: ", colors.Reset,
					)
					input, _ := reader.ReadString('\n')
					query = strings.TrimSpace(input)
				}

				if animeId == -1 {
					err := u.GetAnimeList(u.GetAnimeListParams[types.NativeAnimeList]{
						AnimeList: &animeList,
						Query:     query,
						Limit:     limit,
						Offset:    offset,
						Fields:    detailFields,
					})
					if err != nil {
						fmt.Printf("%v\n****ERROR GETTING ANIMELIST**** ", err)
						fmt.Println()
						query = "" // retry with new query
						continue
					}
				}

				animeId, err = u.FzfAnimeList(u.FzfAnimeListParams{
					AnimeList: &animeList,
					Limit:     limit,
					Offset:    &offset,
				})
				if err != nil {
					if strings.Contains(err.Error(), "130") { // 130 for ESC in FZF
						os.Exit(0)
					}

					fmt.Printf("%v\n****Unexpected Error, Please try again!!****", err)
					fmt.Println()

					query = "" // retry with new query
					continue
				}
				if animeId == -1 {
					continue // offset has changed (fetch the new list)
				}

				break
			}

			detailsIdxs, _ := rootCmd.Flags().GetIntSlice("details")

			detailFields := make([]es.AnimeDetailField, 0)
			detailFields = append(detailFields, *e.DefaultDetailFields()...)
			detailFields = append(detailFields, u.MapIndicesToDetailFields(detailsIdxs)...)

			var animeDetails types.NativeAnimeDetails
			u.GetAnimeDetails(u.GetAnimeDetailsParams[types.NativeAnimeDetails]{
				AnimeDetails: &animeDetails,
				AnimeId:      animeId,
				DetailType:   "custom",
				Fields:       detailFields,
			})

			animeDetailsUI := ui.AnimeDetailsUI{
				Details:      &animeDetails,
				DetailFields: &detailFields,
			}

			app := tview.NewApplication()
			if err := app.SetRoot(animeDetailsUI.CreateLayout(), true).EnableMouse(true).Run(); err != nil {
				panic(err)
			}
		}
	},
}

func init() {

	searchCmd.Example = fmt.Sprintf(`
        ani-cli search "evangelion" -d=1,2,31
        ani-cli search "samurai champloo" -d 1,2,31
        `,
	)

	rootCmd.AddCommand(searchCmd)
}
