package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/PI-Victor/gopushbullet/pkg/auth"
)

// APIToken stores the Pushbullet API Token that is specified by the user
var APIToken string

// LoginCommand asks for the login token and will store it if it will
// authenticate
var LoginCommand = &cobra.Command{
	Use:   "login",
	Short: "login --token <my_generated_access_token>",
	Long: `
Fill in your Pushbullet Access Token and use it to authenticate to the
Pushbullet API.`,
	Run: func(cmd *cobra.Command, args []string) {
		auth.Authenticate(APIToken)
	},
}

// LogoutCommand remove the current API token from the application
var LogoutCommand = &cobra.Command{
	Use:   "logout",
	Short: "logout removes the current used PushBulelt API access token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout")
	},
}

// ListPushes lists pushes from your Pushbullet account via the API
var ListPushes = &cobra.Command{
	Use:   "list-pushes",
	Short: "List your stored pushes from your Pushbullet Account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListPushes")
	},
}

func init() {
	LoginCommand.PersistentFlags().StringVar(&APIToken, "token", "", "Specify your account Access Token")
}
