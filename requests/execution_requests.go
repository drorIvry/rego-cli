package requests

import (
	"net/http"
)

func GetLatestExecution(baseURL string, definitionId string) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodGet,
		baseURL+"/api/v1/task/"+definitionId+"/latest",
		nil,
	)
	
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}


func AbortTask(baseURL string, executionId string) ([]byte, error) {
	responseBytes, err := InvokeRequest(
		http.MethodPost,
		baseURL+"/api/v1/execution/"+executionId+"/abort",
		nil,
	)
	
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
