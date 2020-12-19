package httperr

import (
	"fmt"
	"net/http"
)

type StrMsgResponse struct {
	HttpStatus int
	Message    string
}

func (e StrMsgResponse) Error() string {
	return fmt.Sprintf("[message reponse] http-status:%d message: %s", e.HttpStatus, e.Message)
}

func BadRequest(format string, a ...interface{}) StrMsgResponse {
	return StrMsgResponse{
		HttpStatus: http.StatusBadRequest,
		Message:    fmt.Sprintf(format, a...),
	}
}

func NotFound(format string, a ...interface{}) StrMsgResponse {
	return StrMsgResponse{
		HttpStatus: http.StatusNotFound,
		Message:    fmt.Sprintf(format, a...),
	}
}

func Forbidden(format string, a ...interface{}) StrMsgResponse {
	return StrMsgResponse{
		HttpStatus: http.StatusForbidden,
		Message:    fmt.Sprintf(format, a...),
	}
}
