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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
