package usercommands

import (
	"fmt"
	"os"
	"strings"

	animelist "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/user-commands/anime-list"
	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	pkgE "github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/spf13/cobra"
)

var MeCmd = &cobra.Command{
	Use:   "me",
	Short: "get user specific data",
	Run: func(cmd *cobra.Command, args []string) {
		userDetails := &types.NativeUserDetails{}
		err := utils.GetUserDetails(utils.NativeUserDetailsParams{
			UserDetails: userDetails,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		// show user info
		list := pkgE.UserAnimeListStatuses()
		listStr := make([]string, 0)
		for _, v := range list {
			listStr = append(listStr, string(v))
		}

		for {
			chosenListType, err := utils.FzfUserMenu(listStr, userDetails)
			if err != nil {
				if strings.Contains(err.Error(), "130") { // 130 for ESC in FZF
					return
				}

				fmt.Printf("%v\n****Unexpected Error, Please try again!!****", err)
				fmt.Println()

				//TODO: be more fault tolerant
				os.Exit(1)
			}

            // NOTE: basically treating these commands "run" functions as a utility function
            //  so don't forget to pass all the context required for executing these commands
			if chosenListType == pkgE.ULS_ALL {
				animelist.AllCmd.Run(cmd, args)
			} else if chosenListType == pkgE.ULS_Watching {
				animelist.WatchingCmd.Run(cmd, args)
			} else if chosenListType == pkgE.ULS_PlanToWatch {
				animelist.PlanToWatchCmd.Run(cmd, args)
			} else if chosenListType == pkgE.ULS_OnHold {
				animelist.OnHoldCmd.Run(cmd, args)
			} else if chosenListType == pkgE.ULS_Dropped {
				animelist.DroppedCmd.Run(cmd, args)
			} else if chosenListType == pkgE.ULS_Completed {
				animelist.CompletedCmd.Run(cmd, args)
			}

		}
		// fetch the user's anime list as the choosen status
	},
}

// sort option should be present in me
func init() {
	// option: --sort
	sortOptionsStr := strings.Builder{}
	sortOptionsStr.WriteString("\n\t\t")
	for i, option := range pkgE.UserAnimeListSortOptions() {
		sortOptionsStr.WriteString(fmt.Sprintf("%d => %s", i, option))
		sortOptionsStr.WriteString("\n\t\t")
	}
	MeCmd.PersistentFlags().Int("sort", 1, strings.TrimSpace(fmt.Sprintf(`
        On what basis the list should be sorted

            Available Options: %s
        `,
		sortOptionsStr.String(),
	)))

	MeCmd.AddCommand(LoginCmd)
	MeCmd.AddCommand(animelist.AllCmd)
	MeCmd.AddCommand(animelist.WatchingCmd)
	MeCmd.AddCommand(animelist.PlanToWatchCmd)
	MeCmd.AddCommand(animelist.CompletedCmd)
	MeCmd.AddCommand(animelist.OnHoldCmd)
	MeCmd.AddCommand(animelist.DroppedCmd)
}
