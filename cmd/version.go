/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Long:  `Print the CLI version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: ", rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
