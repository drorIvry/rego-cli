package requests

import (
	"io/ioutil"
	"net/http"
)

func Ping(baseUrl string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		baseUrl+"/ping",
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
