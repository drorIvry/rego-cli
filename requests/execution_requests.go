package requests

import (
	"io/ioutil"
	"net/http"
)

func GetLatestExecution(baseURL string, definitionId string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"/api/v1/task/"+definitionId+"/latest",
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


func AbortTask(baseURL string, executionId string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		baseURL+"/api/v1/execution/"+executionId+"/abort",
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
