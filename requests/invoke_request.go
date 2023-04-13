package requests


import (
	"io"
	"io/ioutil"
	"net/http"
)


func InvokeRequest(
	method string,
	url string,
	requestBody io.Reader,
) ([]byte, error) {
	req, err := http.NewRequest(
		method,
		url,
		requestBody,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}
