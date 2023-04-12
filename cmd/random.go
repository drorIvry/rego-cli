// /*
// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
    
type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	res := getJokeData(url)
	joke := Joke{}
	err := json.Unmarshal(res, &joke)
	if err != nil {
		log.Printf("Error in request: %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("Error in request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "test-cli (demo app)")
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Error in request: %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error in request: %v", err)
	}

	return responseBytes
}
