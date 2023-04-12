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

// abortCmd represents the abort command
var abortCmd = &cobra.Command{
	Use:   "abort [EXECUTION ID]",
	Short: "Abort a running task given its execution ID",
	Long: `Abort a running task given its execution ID For example:

	rego abort 3264462c-6311-46e3-b791-22fac75fffde`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("The executionId argument is required")
		}

		executionId := args[0]
		baseUrl := viper.GetString("baseUrl")
		res, err := requests.AbortTask(baseUrl, executionId)
		if err != nil {
			fmt.Println("Error connecting to rego: ", err)
		} else {
			fmt.Println(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(abortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// abortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// abortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
