package rspx

import (
	"fmt"
	"net/http"
)

type ErrorDetail struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    int    `json:"code"`
}

func (e ErrorDetail) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

const (
	Success           = 0
	System            = 1
	Unauthorized      = 3
	InvalidData       = 5
	Referral          = 6
	User              = 7
	Wallet            = 9
	NotFound          = 8
	Contest           = 10
	PasswordIncorrect = 11
	AccountLocked     = 12
	AccountDisabled   = 13
	AccountDeleted    = 14
	UserAlreadyExists = 15
)

// NewError creates a new ErrorDetail with appropriate HTTP status and message
func NewError(err error, code int, status int) error {
	if status == 0 {
		status = http.StatusInternalServerError
	}

	// Xử lý Message từ err
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return &ErrorDetail{
		Code:    code,
		Status:  status,
		Message: msg,
	}
}
