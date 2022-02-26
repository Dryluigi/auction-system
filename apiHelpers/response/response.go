package response

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildErrorResponse(w http.ResponseWriter, code int, err error, data interface{}) {
	message := formatErrorMessage(code, err)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Status:  code,
		Error:   message,
		Data:    data,
	})
}

func BuildSuccessResponse(w http.ResponseWriter, code int, message string, isSuccess bool, data interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccessResponse{
		Success: isSuccess,
		Status:  code,
		Message: message,
		Data:    data,
	})
}

func formatErrorMessage(status int, err error) string {
	message := "An Error Occured"

	if debugMode, errParse := strconv.ParseBool(os.Getenv("DEBUG_MODE")); errParse == nil && debugMode {
		message = err.Error()
	} else {
		switch status {
		case http.StatusBadRequest:
			message = "Bad Request"
		case http.StatusForbidden:
			message = "Forbidden"
		case http.StatusUnauthorized:
			message = "Unauthorized"
		case http.StatusNotFound:
			message = "Not Found"
		case http.StatusInternalServerError:
			message = "Internal Server Error"
		}
	}

	return message
}
