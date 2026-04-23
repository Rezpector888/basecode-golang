package utils

import (
	"encoding/json"
	"net/http"

	"example.com/common/config"
)

type ResSuccess struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResError struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResPagination struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
}

func ResponseWrite(w http.ResponseWriter, responseCode int, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		config.LogMessage("ERROR", "Failed to encode response: "+err.Error())
	}
}

func ResponseSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	if message == "" {
		message = "Success"
	}

	ResponseWrite(w, statusCode, ResSuccess{
		Status:  true,
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

func ResponseError(w http.ResponseWriter, statusCode int, message string) {

	if len(message) == 0 {
		message = "Internal Server Error"
	}
	if (statusCode) == 0 {
		statusCode = http.StatusInternalServerError
	}

	ResponseWrite(w, statusCode, ResError{
		Status:  false,
		Code:    statusCode,
		Message: message,
	})
}

func ResponsePagination(w http.ResponseWriter, statusCode int, message string, data interface{}, total, limit, offset int) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	if message == "" {
		message = "Success"
	}

	ResponseWrite(w, statusCode, ResPagination{
		Status:  true,
		Code:    statusCode,
		Message: message,
		Data:    data,
		Total:   total,
		Limit:   limit,
		Offset:  offset,
	})
}
