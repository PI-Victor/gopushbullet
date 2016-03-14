package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/PI-Victor/gopushbullet/pkg/auth"
	"github.com/PI-Victor/gopushbullet/pkg/client"
)

// APIToken stores the Pushbullet API Token that is specified by the user
var (
	APIToken     string
	pushesFilter string
)

// LoginCommand asks for the login token and will store it if it will
// authenticate
var LoginCommand = &cobra.Command{
	Use:   "login",
	Short: "login --token <my_generated_access_token>",
	Long: `
Fill in your Pushbullet Access Token and use it to access you account.`,
	Run: func(cmd *cobra.Command, args []string) {
		auth.Authenticate(APIToken)
	},
}

// LogoutCommand remove the current API token from the application
var LogoutCommand = &cobra.Command{
	Use:   "logout",
	Short: "logout removes the current stored user details that are currently used for authentication",
	Run: func(cmd *cobra.Command, args []string) {
		newConfig := client.NewConfig()
		newConfig.PurgeConfig()
		fmt.Println("Your user details have been successfully removed")
	},
}

// ListPushes lists pushes from your Pushbullet account via the API
var ListPushes = &cobra.Command{
	Use:   "list-pushes",
	Short: "List stored pushes from your account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListPushes", pushesFilter)
	},
}

func init() {
	LoginCommand.PersistentFlags().StringVar(&APIToken, "token", "", "Specify your account Access Token")
	ListPushes.PersistentFlags().StringVar(&pushesFilter, "filter", "", "A filter for the returned pushes")
}
