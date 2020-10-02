package errtype

import (
	"fmt"
	"net/http"
)

type ForMsgResponse struct {
	HttpStatus int
	Message    string
}

func (e ForMsgResponse) Error() string {
	return fmt.Sprintf("[message reponse] http-status:%d message: %s", e.HttpStatus, e.Message)
}

func BadRequest(format string, a ...interface{}) ForMsgResponse {
	return ForMsgResponse{
		HttpStatus: http.StatusBadRequest,
		Message:    fmt.Sprintf(format, a...),
	}
}

func NotFound(format string, a ...interface{}) ForMsgResponse {
	return ForMsgResponse{
		HttpStatus: http.StatusNotFound,
		Message:    fmt.Sprintf(format, a...),
	}
}

func Forbidden(format string, a ...interface{}) ForMsgResponse {
	return ForMsgResponse{
		HttpStatus: http.StatusForbidden,
		Message:    fmt.Sprintf(format, a...),
	}
}
