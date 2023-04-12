/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/drorivry/rego-cli/requests"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateFile string
var updateData string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update -d [DATA]",
	Short: "update a task definition",
	Long: `Update a task definition. For example:


	rego update -d '{"id": "3264462c-6311-46e3-b791-22fac75fffde", "image": "hello-world"}'
	rego update -f /path/to/request.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateFile == "" && updateData == "" {
			fmt.Println("No data provided")
			return
		}

		var request io.Reader

		baseUrl := viper.GetString("baseUrl")

		if updateData != "" {
			request = strings.NewReader(updateData)
		} else {
			fileData, err := os.Open(updateFile)
			if err != nil {
				fmt.Println("Error connecting to rego: ", err)
			} else {
				request = fileData
			}
		}

		res, err := requests.UpdateTask(baseUrl, request)
		if err != nil {
			fmt.Println("Error connecting to rego: ", err)
		} else {
			fmt.Println(string(res))
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateFile, "file", "f", "", "A file containing the definition json")
	updateCmd.Flags().StringVarP(&updateData, "data", "d", "", "the definition json")
	rootCmd.AddCommand(updateCmd)
}
