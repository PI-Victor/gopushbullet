package main

import (
	"github.com/spf13/cobra"

	"github.com/PI-Victor/gunner/pkg/cmd"
)

func main() {
	RootCmd.AddCommand(cmd.LoginCommand)
	RootCmd.AddCommand(cmd.LogoutCommand)
	RootCmd.AddCommand(cmd.ListPushes)
	RootCmd.AddCommand(cmd.ListDevices)
	RootCmd.Execute()
}

// RootCmd main command for gunner
var RootCmd = &cobra.Command{
	Use:   "gunner",
	Short: "gunner - a CLI client for Pushbullet",
	Long: `
gunner is a CLI wrapper for Pushbullet. Pushbullet is an application that
allows you to share/push important links, notes, send SMS or mirror
notifications to all your devices. You can mainly use gunner to push important
bookmarks to all of your devices or to get bookmarks saved from your devices or
browser.

You need an account at https://www.pushbullet.com/. Now you need to create an
access token for the Pushbullet API. To do that, you need to go to your Account ->
Settings and then Create Access Token. After you have the access token all you
have to do is gunner login --token <my_generated_access_token>. `,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
