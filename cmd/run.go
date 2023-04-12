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

var image string
var file string
var data string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if image == "" && file == "" && data == "" {
			fmt.Println("No data provided")
			return
		}
		baseUrl := viper.GetString("baseUrl")
		var request io.Reader
		if image != "" {
			body := "{\"image\":\"" + image + "\"}"
			request = strings.NewReader(body)
		} else if data != "" {
			request = strings.NewReader(data)
		} else {
			fileData, err := os.Open(file)
			if err != nil {
				fmt.Println("Error connecting to rego: ", err)
			} else {
				request = fileData
			}
		}

		res, err := requests.RunTask(baseUrl, request)
		if err != nil {
			fmt.Println("Error connecting to rego: ", err)
		} else {
			fmt.Println(string(res))
		}
	},
}

func init() {
	runCmd.Flags().StringVarP(&image, "image", "i", "", "The image to run")
	runCmd.Flags().StringVarP(&file, "file", "f", "", "A file containing the definition json")
	runCmd.Flags().StringVarP(&data, "data", "d", "", "the definition json")
	rootCmd.AddCommand(runCmd)

}
