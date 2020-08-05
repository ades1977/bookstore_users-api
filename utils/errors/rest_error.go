package errors

import "net/http"

type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []string `json:"data"`
}



func NewBedrequest(message string)*RestErr{
	return &RestErr{
		Status: http.StatusBadRequest,
		Message: message,
		Data: []string{ "bed_request"},
	}
}

func NewNotFound(message string)*RestErr{
	return &RestErr{
		Status: http.StatusNotFound,
		Message:message,
		Data: []string{"bed_not_found"},
	}
}

func NewInternalServerError(message string)*RestErr{
	return &RestErr{
		Status: http.StatusInternalServerError,
		Message: message,
		Data: []string{"internal_server_error"},
	}
}

func NewSaveDBSuccess(message string) *RestErr{
	return &RestErr{
		Status: http.StatusOK,
		Message: message,
		Data: []string{"server_status_OK"},
	}
}