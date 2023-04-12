package cmd

import (
	"fmt"

	"github.com/drorivry/rego-cli/requests"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping to check the connection to the rego server",
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl := viper.GetString("baseUrl")
		res, err := requests.Ping(baseUrl)
		if err != nil {
			fmt.Println("Error connecting to rego: ", err)
		} else {
			fmt.Println(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
