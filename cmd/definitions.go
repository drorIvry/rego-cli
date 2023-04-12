package cmd

import (
	"fmt"

	"github.com/drorivry/rego-cli/requests"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// definitionsCmd represents the definitions command
var definitionsCmd = &cobra.Command{
	Use:   "definitions",
	Short: "Get a list if all the task definitions",
	Long: `Get a list if all the task definitions For example:

	rego definitions
`,
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl := viper.GetString("baseUrl")
		res, err := requests.GetAllTasks(baseUrl)
		if err != nil {
			fmt.Println("Error connecting to rego: ", err)
		} else {
			fmt.Println(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(definitionsCmd)
}
