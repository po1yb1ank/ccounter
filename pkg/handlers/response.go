package handlers

import "net/http"

type CounterResponse struct {
	Status       int    `json:"status"`
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`

	Key   string `json:"key,omitempty"`
	Value int64  `json:"value,omitempty"`
}

func NewOkGetCounterResponse(key string, value int64) *CounterResponse {
	return &CounterResponse{
		Status:  http.StatusOK,
		Message: "Successful get key value action",
		Key:     key,
		Value:   value,
	}
}

func NewErrCounterResponse(errorMessage string) *CounterResponse {
	return &CounterResponse{
		Status:       http.StatusInternalServerError,
		ErrorMessage: errorMessage,
	}
}

func NewErrWSConnectionResponse() *CounterResponse {
	return &CounterResponse{
		Status:       http.StatusBadRequest,
		ErrorMessage: ErrorWSConn.Error(),
	}
}
