package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var baseUrl string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rego-cli [deploy, update, abort, rerun, ]",
	Short: "CLI tool for rego.",
	Long:  `Work seamlessly with GitHub from the command line.`,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is  $HOME/.rego/.rego-cli.yaml)")
	rootCmd.PersistentFlags().StringVar(&baseUrl, "baseUrl", "http://localhost:4004", "The rego server URL (default is 'http://localhost:4004')")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.BindPFlag("baseUrl", rootCmd.PersistentFlags().Lookup("baseUrl"))

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra-example" (without extension).
		viper.SetConfigName(".rego-cli")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home + "/.rego")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error loading config file:", err)
	}
}
