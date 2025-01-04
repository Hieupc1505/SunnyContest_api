package response

import (
	"SunnyContest/global"
	"fmt"
	"go.uber.org/zap"
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

func WrapError(err error, category, key string) error {
	if global.Config.Server.Mode == "production" {
		global.Logger.Error(err.Error(), zap.String("Error", "error"))
	}
	if global.Config.Server.Mode == "dev" {
		fmt.Println(err)
	}
	return GetError(category, key)
}

type ErrorCategory struct {
	Prefix int
	Errors map[string]ErrorDetail
}

// GetError: Tạo mã lỗi dựa trên danh mục và lỗi
func GetError(category, key string) error {
	if cat, ok := ErrorCatalog[category]; ok {
		if key == "" {
			return ErrorDetail{
				Code: cat.Prefix,
			}
		}
		if err, ok := cat.Errors[key]; ok {
			// Ghép mã lỗi: Prefix + Code
			err.Code = cat.Prefix
			return err
		}
	}
	return ErrorDetail{
		Code:    ErrorCatalog["System"].Prefix,
		Message: "Faile to get error with",
	}
}

// ErrorCatalog: tổ chức lỗi thành từng nhóm
var ErrorCatalog = map[string]ErrorCategory{
	"System": {
		Prefix: 1,
		Errors: map[string]ErrorDetail{
			"SystemError": {Message: "system_error", Status: http.StatusInternalServerError},
		},
	},
	"Unauthorized": {
		Prefix: 3,
		Errors: map[string]ErrorDetail{
			"SessionExpired": {Message: "unauthorized.session_expired", Status: http.StatusUnauthorized},
			"InvalidToken":   {Message: "unauthorized.invalid_token", Status: http.StatusUnauthorized},
		},
	},
	"InvalidData": {
		Prefix: 5,
		Errors: map[string]ErrorDetail{
			"General":         {Message: "invalid_data", Status: http.StatusBadRequest},
			"ID":              {Message: "invalid_data.id", Status: http.StatusBadRequest},
			"Name":            {Message: "invalid_data.name", Status: http.StatusBadRequest},
			"QuestionLevel":   {Message: "invalid_data.question.level", Status: http.StatusBadRequest},
			"QuestionAnswer":  {Message: "invalid_data.question.answer", Status: http.StatusBadRequest},
			"AnswerType":      {Message: "invalid_data.question.answer_type", Status: http.StatusBadRequest},
			"QuestionType":    {Message: "invalid_data.question.type", Status: http.StatusBadRequest},
			"SubjectID":       {Message: "invalid_data.subject_id", Status: http.StatusBadRequest},
			"TimeExam":        {Message: "invalid_data.contest.time_exam", Status: http.StatusBadRequest},
			"SubjectName":     {Message: "invalid_data.contest.subject_name", Status: http.StatusBadRequest},
			"NotEnoughAmount": {Message: "invalid_data.not_enough_amount", Status: http.StatusBadRequest},
			"ContestNotFound": {Message: "invalid_data.contest.not_found", Status: http.StatusBadRequest},
			"Amount":          {Message: "invalid_data.amount", Status: http.StatusBadRequest},
			"Receiver":        {Message: "invalid_data.receiver", Status: http.StatusBadRequest},
			"TransferToSelf":  {Message: "invalid_data.transfer_to_self", Status: http.StatusBadRequest},
			"ContestGameID":   {Message: "invalid_data.contest.game_id", Status: http.StatusBadRequest},
		},
	},
	"Contest": {
		Prefix: 12,
		Errors: map[string]ErrorDetail{
			"Running":     {Message: "contest.running", Status: http.StatusBadRequest},
			"Finished":    {Message: "contest.finished", Status: http.StatusBadRequest},
			"LiveAlready": {Message: "contest.live_already", Status: http.StatusBadRequest},
			"LiveSubmit":  {Message: "contest.live_submit_already", Status: http.StatusBadRequest},
		},
	},
	"Wallet": {
		Prefix: 9,
		Errors: map[string]ErrorDetail{
			"NotFound":  {Message: "wallet.not_found", Status: http.StatusBadRequest},
			"NotEnough": {Message: "wallet.balance.not_enough", Status: http.StatusBadRequest},
		},
	},
	"Referral": {
		Prefix: 6,
		Errors: map[string]ErrorDetail{
			"Claimed":     {Message: "referral.claimed", Status: http.StatusBadRequest},
			"NotFinished": {Message: "referral.not_finished", Status: http.StatusBadRequest},
		},
	},
	"User": {
		Prefix: 7,
		Errors: map[string]ErrorDetail{
			"NotFound": {Message: "user.not_found", Status: http.StatusNotFound},
		},
	},
	"Data": {
		Prefix: 8,
		Errors: map[string]ErrorDetail{
			"NotFound": {Message: "data.not_found", Status: http.StatusNotFound},
		},
	},
	"PasswordIncorrect": {
		Prefix: 11,
		Errors: map[string]ErrorDetail{},
	},
	"AccountLocked": {
		Prefix: 12,
		Errors: map[string]ErrorDetail{},
	},
	"AccountDisabled": {
		Prefix: 13,
		Errors: map[string]ErrorDetail{},
	},
	"AccountDeleted": {
		Prefix: 14,
		Errors: map[string]ErrorDetail{},
	},
	"UserAlreadyExists": {
		Prefix: 15,
		Errors: map[string]ErrorDetail{},
	},
}
