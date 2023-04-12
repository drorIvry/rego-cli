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