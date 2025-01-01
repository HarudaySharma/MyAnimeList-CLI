package animelist

import (
	"fmt"
	"os"
	"strings"

	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/ui"
	u "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	pkgE "github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/spf13/cobra"
)

var PlanToWatchCmd = &cobra.Command{
	Use:   "watching",
	Short: "lists all the user's anime",
	Run: func(cmd *cobra.Command, args []string) {
		listType := pkgE.ULS_PlanToWatch

		offset := 0
		limit, err := cmd.Flags().GetInt("list-size")
		if err != nil {
			limit = 10 // don't panic
		}

		if limit > es.User_Max_List_Size {
			limit = es.User_Default_List_Size
		}

		sortBy, err := cmd.Flags().GetInt("sort")
		if err != nil {
			sortBy = 1 // anime_score
		}

		var animeList types.NativeUserAnimeList
		var animeId int = -1 // will send a request to server every time it is "-1"

		detailFields := make([]es.AnimeDetailField, 0)
		detailFields = append(detailFields, *e.PreviewDetailFields()...)
		detailFields = append(detailFields, es.Title, es.StartSeason)

		for {
			for {

				if animeId == -1 {

					err := u.GetUserAnimeList(u.GetUserAnimeListParams[types.NativeUserAnimeList]{
						AnimeList: &animeList,
						ListType:  listType,
						Limit:     limit,
						Offset:    offset,
						Sort:      sortBy,
						Fields:    detailFields,
					})
					if err != nil {
						fmt.Printf("%v\n****ERROR IN USER ANIME LIST CMD**** ", err)
						fmt.Println()
						//TODO: be more fault tolerant
						os.Exit(1)
					}
				}

				animeId, err = u.FzfUserAnimeList(u.FzfUserAnimeListParams{
					AnimeList: &animeList,
					Limit:     limit,
					Offset:    &offset,
				})
				if err != nil {
					if strings.Contains(err.Error(), "130") { // 130 for ESC in FZF
						return
					}

					fmt.Printf("%v\n****Unexpected Error, Please try again!!****", err)
					fmt.Println()

					//TODO: be more fault tolerant
					os.Exit(1)

				}
				if animeId == -1 {
					continue // offset has changed (fetch the new list)
				}

				break
			}

			detailsIdxs, _ := cmd.Flags().GetIntSlice("details")

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

			app := ui.NewApplication(&animeDetailsUI)
			if err := app.Run(); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	// option: --sort
	sortOptionsStr := strings.Builder{}
	sortOptionsStr.WriteString("\n\t\t")
	for i, option := range pkgE.UserAnimeListSortOptions() {
		sortOptionsStr.WriteString(fmt.Sprintf("%d => %s", i, option))
		sortOptionsStr.WriteString("\n\t\t")
	}
	PlanToWatchCmd.PersistentFlags().Int("sort", 1, strings.TrimSpace(fmt.Sprintf(`
        On what basis the list should be sorted

            Available Options: %s
        `,
		sortOptionsStr.String(),
	)))

}
