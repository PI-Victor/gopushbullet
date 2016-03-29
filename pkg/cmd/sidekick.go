package cmd

import (
	"github.com/spf13/cobra"

	"github.com/PI-Victor/gunner/pkg/client"
)

// SyncCommand downloads the latest stored pushes
var SyncCommand = &cobra.Command{
	Use:   "sync",
	Short: "sidekick sync",
	Long:  `Syncronize pushes stored remotely with the ones stored locally`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SyncPushes()
	},
}
