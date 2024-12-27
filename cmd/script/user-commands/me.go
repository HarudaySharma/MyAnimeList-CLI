package usercommands

import (
	"fmt"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
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

		showUserInfo(userDetails)

	},
}

func showUserInfo(userDetails *types.NativeUserDetails ) {
    utils.FzfUserDetails(userDetails)
}

func init() {
	MeCmd.AddCommand(loginCmd)
}
