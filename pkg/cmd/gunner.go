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
login --google user@gmail.com
login --facebook user@facebook.com
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

// PushNote pushes a message/notification to a specified device or to all
var PushNote = &cobra.Command{
	Use:   "push-note",
	Short: "Push ephemeral notifications from the CLI to your devices",
	Run: func(cmd *cobra.Command, args []string) {
		client.PushNote()
	},
}

// PushSMS sends a SMS message to a specific number from your device
var PushSMS = &cobra.Command{
	Use:   "sms",
	Short: "Send SMS messages on behalf of your device to a specified number",
	Run: func(cmd *cobra.Command, args []string) {
		client.PushSMS()
	},
}

// ListDevices lists the devices that are attached to your Pushbullet account
var ListDevices = &cobra.Command{
	Use:   "list-devices",
	Short: "List devices that are attached to your account",
	Run: func(cmd *cobra.Command, args []string) {
		client.ListDevices()
	},
}

func init() {
	LoginCommand.PersistentFlags().StringVar(&APIToken, "token", "", "Specify your account Access Token")
	LoginCommand.PersistentFlags().StringVar(&APIToken, "oauth", "", "Specify an OAuth provider")
	ListPushes.PersistentFlags().StringVar(&pushesFilter, "filter", "", "A filter for the returned pushes")
	ListPushes.PersistentFlags().StringVar(&deleted, "deleted", "off", "Display deleted pushes by turning the flag on")
}
