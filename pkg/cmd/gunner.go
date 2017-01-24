package cmd

import (
	"github.com/spf13/cobra"

	"github.com/PI-Victor/gunner/pkg/client"
)

// APIToken stores the Pushbullet API Token that is specified by the user
var (
	APIToken     string
	pushesFilter string
	deleted      string
)

// LoginCommand asks for the login token and will store it if it will
// authenticate
var LoginCommand = &cobra.Command{
	Use:   "login",
	Short: "Login with a generated token or an oauth provider",
	Long: `
Fill in your Pushbullet Access Token or specify an oauth provider and use it to
access you account.
login --token <my_generated_access_token>
`,
	Run: func(cmd *cobra.Command, args []string) {
		user := client.NewUser()
		user.Token = APIToken
		user.Authenticate()
	},
}

// LogoutCommand remove the current API token from the application
var LogoutCommand = &cobra.Command{
	Use:   "logout",
	Short: `Removes the current stored user details that are used for authentication`,
	Run: func(cmd *cobra.Command, args []string) {
		newConfig := client.NewConfig()
		newConfig.Logout()
	},
}

// ListPushes lists pushes from your Pushbullet account via the API
var ListPushes = &cobra.Command{
	Use:   "list-pushes",
	Short: "List stored pushes from your account",
	Run: func(cmd *cobra.Command, args []string) {
		client.ListPushes()
	},
}

func init() {
	LoginCommand.PersistentFlags().StringVar(&APIToken, "token", "", "Specify your account Access Token")
	ListPushes.PersistentFlags().StringVar(&pushesFilter, "filter", "", "A filter for the returned pushes")
	ListPushes.PersistentFlags().StringVar(&deleted, "deleted", "off", `Pushes that have been delete are
		off by default, you can enable them by turning the flag on`)
}
