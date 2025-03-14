package script

import (
	"fmt"
	"os"
	"strings"

	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/ui"
	u "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/spf13/cobra"
)

var rankingCmd = &cobra.Command{
	Use:   "ranking",
	Short: "get list of anime based on ranking",
	Run: func(cmd *cobra.Command, args []string) {

		offset := 0
		limit, err := rootCmd.Flags().GetInt("list-size")
		if err != nil {
			limit = 10 // don't panic
		}

		if limit > es.MAX_RANKING_SEARCH_LIST_SIZE {
			limit = es.MAX_RANKING_SEARCH_LIST_SIZE
		}

		rankingType, err := cmd.Flags().GetString("ranking-type")
		if err != nil {
			rankingType = string(es.RankingAiring)
		}

		var animeList types.NativeAnimeRanking
		var listNode *types.AnimeRankingDataNode
		var animeId int = -1 // will send a request to server every time it is "-1"

		detailFields := make([]es.AnimeDetailField, 0)
		detailFields = append(detailFields, *e.PreviewDetailFields()...)
		detailFields = append(detailFields, es.Title)

		for {
			for {

				if animeId == -1 {
					err := u.GetAnimeRanking(u.GetAnimeRankingParams[types.NativeAnimeRanking]{
						AnimeList:   &animeList,
						RankingType: rankingType,
						Limit:       limit,
						Offset:      offset,
						Fields:      *e.PreviewDetailFields(),
					})
					if err != nil {
						fmt.Printf("%v\n****ERROR IN RANKING ANIME****", err)
						fmt.Println()
						//TODO: be more fault tolerant
						os.Exit(1)
					}
				}

				animeId, listNode, err = u.FzfRankingAnimeList(u.FzfRankingAnimeListParams{
					AnimeList:   &animeList,
					Limit:       limit,
					Offset:      &offset,
					RankingType: rankingType,
				})
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
				ListNode:     listNode,
			}

			app := ui.NewApplication(&animeDetailsUI)
			if err := app.Run(); err != nil {
				panic(err)
			}
		}
	},
}

func init() {

	animeRankingTypesStr := strings.Builder{}
	animeRankingTypesStr.WriteString("\n\t\t")
	for _, t := range es.AnimeRanking() {
		animeRankingTypesStr.WriteString(fmt.Sprintf("%s", t))
		animeRankingTypesStr.WriteString("\n\t\t")
	}

	// --ranking-type
	rankingCmd.PersistentFlags().StringP("ranking-type", "t", string(es.AnimeRanking()[0]), strings.TrimSpace(fmt.Sprintf(`
        Available Options:
        %s
        `,
		animeRankingTypesStr.String())))

	rankingCmd.Example = fmt.Sprintf(`
        ani-cli ranking -t %s
        ani-cli ranking --type %s
        `,
		es.AnimeRanking()[0],
		es.AnimeRanking()[0],
	)

	rootCmd.AddCommand(rankingCmd)
}
