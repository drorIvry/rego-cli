package requests

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetLatestExecution(baseURL string, definitionId string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		baseURL,
		nil,
	)

	if err != nil {
		log.Printf("Error in request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
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
