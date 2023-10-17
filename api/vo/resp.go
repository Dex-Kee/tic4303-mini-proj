package vo

import "net/http"

type ResponseVO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResp(data any) *ResponseVO {
	return &ResponseVO{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

func ErrorResp(code int, message string) *ResponseVO {
	return &ResponseVO{
		Code:    code,
		Message: message,
		Data:    "",
	}
}

func BadRequestResp(message string) *ResponseVO {
	return &ResponseVO{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    "",
	}
}

func UnauthorizedResp(message string) *ResponseVO {
	return &ResponseVO{
		Code:    http.StatusUnauthorized,
		Message: message,
		Data:    "",
	}
}

func ForbiddenResp(message string) *ResponseVO {
	return &ResponseVO{
		Code:    http.StatusForbidden,
		Message: message,
		Data:    "",
	}
}
