package response

type Response struct {
	Errors interface{} `json:"e"`
	Data   interface{} `json:"d,omitempty"`
}

type MsgError struct {
	Msg string `json:"msg,omitempty"`
}

func SuccessResponseData(data interface{}) Response {
	return Response{
		Errors: 0,
		Data:   data,
	}
}

func ErrorResponseData(err ErrorDetail) Response {
	return Response{
		Errors: err.Code,
		Data:   MsgError{Msg: err.Message},
	}
}

func HandleErrorResponse(err error) (Response, int) {
	switch e := err.(type) {
	case ErrorDetail:
		return ErrorResponseData(e), e.Status
	default:
		return ErrorResponseData(ErrorDetail{
			Code:    500,
			Message: GetError("System", "SystemError").Error(),
		}), 500
	}
}
