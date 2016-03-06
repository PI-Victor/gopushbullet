package main

import (
	"github.com/PI-Victor/gopushbullet/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	RootCmd.AddCommand(cmd.LoginCommand)
	RootCmd.AddCommand(cmd.LogoutCommand)
	RootCmd.Execute()
}

// RootCmd main command for monito
var RootCmd = &cobra.Command{
	Use:   "gopush",
	Short: "gopush - a CLI client for Pushbullet",
	Long: `
gopush is a CLI wrapper for Pushbullet on linux. Pushbullet is an application
that allows you to share/push important things to all your devices. You can
mainly use gopush to push important bookmarks to all of your devices or to get
bookmarks saved from your devices or browser.

First you need to create an account on https://www.pushbullet.com/.  Now you
need to create an access token for the Pushbullet API. To do that, you need to
go to your Account -> Settings and then Create Access Token. After you have the
access token all you have to do is gopush login --token
<my_generated_access_token>. `,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
