package script

import (
	"github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/daemon"
	"github.com/spf13/cobra"
)

var stopDaemon = &cobra.Command{
    Use: "stop_daemon",
    Short: "Stops the mal-cli daemon process running in the background",
    Run: func(cmd *cobra.Command, args []string) {
        daemon.StopDaemon();
    },
}

func init() {
    rootCmd.AddCommand(stopDaemon)
}
