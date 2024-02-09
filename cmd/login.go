package cmd

import (
	"fmt"

	"github.com/drorivry/rego-cli/requests"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to rego server",
	Run: func(cmd *cobra.Command, args []string) {		
		// get the api key from the user
		fmt.Print("Enter your API key: ")
		var apiKey string
		fmt.Scanln(&apiKey)
		
		// store the API key to the config file
		baseUrl := viper.GetString("baseUrl")
		_, err := requests.LoginPing(baseUrl, apiKey)
		
		if err != nil {
			fmt.Println("Error connecting to rego: ", err)
		} else {
			viper.Set("apiKey", apiKey)
			viper.WriteConfig()
			fmt.Println("Logged in successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
