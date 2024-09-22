package script

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mal-cli",
	Short: "Search about anime from terminal",
	Long:  `Access MyAnimList api from terminal`,
	Run: func(cmd *cobra.Command, args []string) {
        searchCmd.Run(cmd, args);
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
}
