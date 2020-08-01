package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status    int    `json:"status"`
	Error   string `json:"error"`
}

func NewBedrequest(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bed_request",
	}
}

func NewNotFound(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "bed_not_found",
	}
}