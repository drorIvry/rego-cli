package requests

import (
	"io"
	"net/http"
)

func RerunTask(baseURL string, definitionId string) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodPost,
		baseURL+"/api/v1/task/"+definitionId+"/rerun",
		nil,
	)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func GetAllTasks(baseURL string) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodGet,
		baseURL+"/api/v1/task/",
		nil,
	)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func DeleteTask(baseURL string, definitionId string) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodDelete,
		baseURL+"/api/v1/task/"+definitionId,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func RunTask(baseURL string, requestBody io.Reader) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodPost,
		baseURL+"/api/v1/task",
		requestBody,
	)

	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func UpdateTask(baseURL string, requestBody io.Reader) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodPut,
		baseURL+"/api/v1/task",
		requestBody,
	)
	
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
