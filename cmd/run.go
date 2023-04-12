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
	Use:   "run -i [IMAGE]",
	Short: "run a new task",
	Long: `Run a new task. For example:

	rego run -i hello-world
	rego run -d '{"image": "hello-world"}'
	rego run -f /path/to/request.json`,
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
