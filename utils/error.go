package utils

const (
	ErrTypeInternalServer = "InternalServerError"
	ErrMsgInternalServer  = "内部サーバーエラーが発生しました。しばらくしてから再度アクセスしてください。"
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
