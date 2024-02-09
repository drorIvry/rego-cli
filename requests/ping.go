package requests

import (
	"net/http"
)

func Ping(baseUrl string) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodGet,
		baseUrl+"/ping",
		nil,
	)
	
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func LoginPing(baseUrl string, apiKey string) ([]byte, error) {
	responseBytes, err := InvokeRequestWithApiKey(
		http.MethodGet,
		baseUrl+"/ping",
		nil,
		apiKey,
	)
	
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
