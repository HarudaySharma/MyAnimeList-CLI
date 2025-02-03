package usercommands

import (
	"fmt"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	pkgU "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "ask's user to give permission to access their mal data",
	Run: func(cmd *cobra.Command, args []string) {
		const mal_auth_url = "https://myanimelist.net/v1/oauth2/authorize"

		url := fmt.Sprintf("%s?response_type=code&client_id=%s&code_challenge=%s&state=Request123&code_challenge_method=plain",
			mal_auth_url,
			config.C.MalClientId,
			config.C.MalCodeChallenge,
		)

		fmt.Println("Redirecting to your browser for authorization...")

		err := pkgU.OpenURL(url)
		if err != nil {
			fmt.Println("Failed to open browser:", err)
			fmt.Println("Please manually visit the URL below to authorize:")
			fmt.Println(url)
		}

		return
	},
}
