package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)


func InvokeRequest(
	method string,
	url string,
	requestBody io.Reader,
) ([]byte, error) {

	apiKey := viper.GetString("apiKey")
	if apiKey == "" {
		fmt.Println("You are not logged in. Please run 'rego login' to log in.")
		return nil, nil
	}

	req, err := http.NewRequest(
		method,
		url,
		requestBody,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Rego-Api-Key", apiKey)
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}


func InvokeRequestWithApiKey(
	method string,
	url string,
	requestBody io.Reader,
	apiKey string,
) ([]byte, error) {
	if apiKey == "" {
		fmt.Println("You are not logged in. Please run 'rego login' to log in.")
		return nil, nil
	}

	req, err := http.NewRequest(
		method,
		url,
		requestBody,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Rego-Api-Key", apiKey)
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
