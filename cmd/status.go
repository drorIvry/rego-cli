/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/drorivry/rego-cli/requests"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status [DEFINITION ID]",
	Short: "Get the latest execution status of a given definition ID",
	Long: `Get the latest execution status of a given definition ID For example:

	rego status 3264462c-6311-46e3-b791-22fac75fffde
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("The definitionId argument is required")
		}

		definitionId := args[0]
		baseUrl := viper.GetString("baseUrl")
		res, err := requests.GetLatestExecution(baseUrl, definitionId)
		if err != nil {
			fmt.Println("Error connecting to rego: ", err)
		} else {
			fmt.Println(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
