package rspx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Errors interface{} `json:"e"`
	Data   interface{} `json:"d,omitempty"`
}

type MsgError struct {
	Msg string `json:"msg,omitempty"`
}

func NewSucessResponse(data interface{}) Response {
	return Response{
		Errors: Success,
		Data:   data,
	}
}

func ErrorResponseData(err *ErrorDetail) Response {
	return Response{
		Errors: err.Code,
		Data:   MsgError{Msg: err.Message},
	}
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewSucessResponse(data))
}

func ErrorResponse(ctx *gin.Context, err error) {
	var e *ErrorDetail
	var data Response
	status := http.StatusInternalServerError
	switch {
	case errors.As(err, &e):
		data = ErrorResponseData(e)
		status = e.Status
	default:
		data = ErrorResponseData(&ErrorDetail{
			Code:    System,
			Message: err.Error(),
		})
	}
	ctx.JSON(status, data)
}
