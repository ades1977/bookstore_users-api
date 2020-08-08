package errors

import "net/http"

type ErrDetail struct {
	ErrorCode 	int `json:"error_code"`
	ErrorMsg	string `json:"error_msg"`
}

type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    ErrDetail `json:"data"`
}



func NewBedrequest(message string)*RestErr{
	return &RestErr{
		Status: http.StatusBadRequest,
		Message: message,
		Data : ErrDetail{
			ErrorCode: http.StatusBadRequest,
			ErrorMsg:  "bed_request",
		},
	}
}

func NewNotFound(message string)*RestErr{
	return &RestErr{
		Status: http.StatusNotFound,
		Message:message,
		Data : ErrDetail{
			ErrorCode: http.StatusNotFound,
			ErrorMsg:  "status_not_found",
		},
	}
}

func NewInternalServerError(message string)*RestErr{
	return &RestErr{
		Status: http.StatusInternalServerError,
		Message: message,
		Data : ErrDetail{
			ErrorCode: http.StatusInternalServerError,
			ErrorMsg:  "internal_server_error",
		},
	}
}

func NewSaveDBSuccess(message string) *RestErr{
	return &RestErr{
		Status: http.StatusOK,
		Message: message,
		Data : ErrDetail{
			ErrorCode: http.StatusOK,
			ErrorMsg:  "http_status_ok",
		},
	}
}

