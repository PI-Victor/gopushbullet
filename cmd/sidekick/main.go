package main

import (
	"github.com/spf13/cobra"

	"github.com/PI-Victor/gunner/pkg/cmd"
)

func main() {
	RootCmd.AddCommand(cmd.SyncCommand)
	RootCmd.Execute()
}

// RootCmd main command for sidekick
var RootCmd = &cobra.Command{
	Use:   "sidekick",
	Short: "sidekick - an organizer for gunner",
	Long: `
sidekick is the organizer of your locally stored data from gunner. It will sync
data locally and will offer extended functionality like allowing you to create local groups for sorting your links.
`,
}
