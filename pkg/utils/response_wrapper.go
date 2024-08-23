package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code       int    `json:"code"`
	Error      string `json:"error"`
	StackTrace string `json:"stacktrace,omitempty"`
}

func JSONSuccess(w http.ResponseWriter, r *http.Request, code int, message string, data interface{}) {
	response := SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	writeJSON(w, code, response)
}

func JSONError(w http.ResponseWriter, r *http.Request, code int, errMsg string, stacktrace string) {
	response := ErrorResponse{
		Code:       code,
		Error:      errMsg,
		StackTrace: stacktrace,
	}

	writeJSON(w, code, response)
}

func writeJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
