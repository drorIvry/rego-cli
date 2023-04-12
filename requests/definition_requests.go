package requests

import (
	"io"
	"io/ioutil"
	"net/http"
)

func RerunTask(baseURL string, definitionId string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		baseURL+"/api/v1/task/"+definitionId+"/rerun",
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

func GetAllTasks(baseURL string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"/api/v1/task/",
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

func DeleteTask(baseURL string, definitionId string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		baseURL+"/api/v1/task/"+definitionId,
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

func RunTask(baseURL string, requestBody io.Reader) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		baseURL+"/api/v1/task/",
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

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
