package utils

import (
	"fmt"
	"net/http"
)

const (
	ErrTypeInternalServer = "InternalServerError"
	ErrMsgInternalServer  = "内部サーバーエラーが発生しました。しばらくしてから再度アクセスしてください。"

	ErrTypeNotfound = "NotfoundError"
	ErrMsgNotfound  = "存在しないデータです"
)

type AppErr struct {
	Type     string
	Code     int
	Message  string
	Internal error
}

func (ce *AppErr) Error() string {
	return ce.Message
}

func NewAppErr(errType, message string, code int, internal error) *AppErr {
	return &AppErr{
		Type:     errType,
		Code:     code,
		Message:  message,
		Internal: internal,
	}
}

func NewAppValidateByErr(cause string) *AppErr {
	return &AppErr{
		Type:    "ValidationError",
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("%sが適切な入力ではないです", cause),
	}
}
