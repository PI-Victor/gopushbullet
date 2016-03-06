package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// LoginCommand asks for the login token and will store it if it will
// authenticate
var LoginCommand = &cobra.Command{
	Use:   "login",
	Short: "login to PushBullet API",
	Long:  "Fill in your PushBullet Token to login to PushBullet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login")
	},
}

// LogoutCommand remove the current API token from the program
var LogoutCommand = &cobra.Command{
	Use:   "logout",
	Short: "logout removes the current used PushBulelt API access token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout")
	},
}

// func init() {
// 	StartCommand.PersistentFlags().StringVar(, "config", "", "Specify a configuration file")
// }
