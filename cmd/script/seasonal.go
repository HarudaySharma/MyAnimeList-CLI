package script

import (
	"fmt"
	"os"
	"strings"
	"time"

	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/ui"
	u "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var seasonalCmd = &cobra.Command{
	Use:   "seasonal",
	Short: "get the seasonal anime list",
	Run: func(cmd *cobra.Command, args []string) {

		offset := 0
		limit, err := cmd.Flags().GetInt("l")
		if err != nil {
			limit = 10 // don't panic
		}

		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			year = time.Now().Year() // set to currentYear
		}

		season, err := cmd.Flags().GetString("season")
		if err != nil {
			season = string(u.CurrentAnimeSeason())
		}

		sortBy, err := cmd.Flags().GetInt("sort")
		if err != nil {
			sortBy = 0
		}

		var animeList *types.NativeAnimeList
		var animeId int = -1 // will send a request to server every time it is "-1"
		for {
			for {

				if animeId == -1 {
					var seasonalAnimeList types.NativeSeasonalAnime

					err := u.GetSeasonalAnime(u.GetSeasonalAnimeParams[types.NativeSeasonalAnime]{
						AnimeList: &seasonalAnimeList,
						Year:      year,
						Season:    season,
						Limit:     limit,
						Offset:    offset,
						Sort:      sortBy,
						Fields: []es.AnimeDetailField{
							es.StartSeason,
						},
					})
					if err != nil {
						fmt.Printf("%v\n****ERROR IN SEASONAL ANIME**** ", err)
						fmt.Println()
						//TODO: be more fault tolerant
						os.Exit(1)
					}

					animeList = u.SeasonalToNativeAnimeList(&seasonalAnimeList)
				}

				animeId, err = u.FzfAnimeList(animeList, limit, &offset)
				if err != nil {
					if strings.Contains(err.Error(), "130") { // 130 for ESC in FZF
						os.Exit(0)
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

			detailsIdxs, _ := cmd.Flags().GetIntSlice("d")

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

	// option: --year
	currentYear := time.Now().Year()
	seasonalCmd.PersistentFlags().Int("year", currentYear, strings.TrimSpace(fmt.Sprintf(`
        Specify which Year's seasonal anime you want

        Default Value: Current Year
        `,
	)))

	// option: --season
	availableSeasonsStr := strings.Builder{}
	availableSeasonsStr.WriteString("\n\t\t")
	for _, season := range es.AnimeSeasons() {
		availableSeasonsStr.WriteString(fmt.Sprintf("%s", season))
		availableSeasonsStr.WriteString("\n\t\t")
	}

	// get the current season
	currentSeason := u.CurrentAnimeSeason()
	seasonalCmd.PersistentFlags().String("season", string(currentSeason), strings.TrimSpace(fmt.Sprintf(`
        Specify which season's anime you want
            Available seasons: %s

        Default Value: Current Season
        `,
		availableSeasonsStr.String(),
	)))

	// option: --sort
	sortOptionsStr := strings.Builder{}
	sortOptionsStr.WriteString("\n\t\t")
	for i, option := range es.SortOptions() {
		sortOptionsStr.WriteString(fmt.Sprintf("%d => %s", i, option))
		sortOptionsStr.WriteString("\n\t\t")
	}
	seasonalCmd.PersistentFlags().Int("sort", 0, strings.TrimSpace(fmt.Sprintf(`
        On what basis the list should be sorted

            Available Options: %s
        `,
		sortOptionsStr.String(),
	)))

	// option: --l
	seasonalCmd.PersistentFlags().Int("l", 10, strings.TrimSpace(fmt.Sprintf(`
        Specify length of anime list { 1 - 100 }
        `,
	)))

	// option: --d
	availableOptionsStr := strings.Builder{}
	availableOptionsStr.WriteString("\n\t\t")
	for i, option := range es.EveryDetailField() {
		availableOptionsStr.WriteString(fmt.Sprintf("%d => %s", i, option))
		availableOptionsStr.WriteString("\n\t\t")
	}

	seasonalCmd.PersistentFlags().IntSlice("d", []int{}, strings.TrimSpace(fmt.Sprintf(`
        Specify which anime detail you want

            Available Options: %s
            Usage:
                ani-cli search "anime title" --d=1,2,31
                ani-cli search "anime title" --d 1,2,31

        `,
		availableOptionsStr.String(),
	)))

	rootCmd.AddCommand(seasonalCmd)
}
