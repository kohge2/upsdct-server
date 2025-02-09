package response

type errorResponse struct {
	Type    string `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Error errorResponse `json:"error"`
}

func NewErrorResponse(errType, message string, code int) ErrorResponse {
	return ErrorResponse{
		Error: errorResponse{
			Type:    errType,
			Code:    code,
			Message: message,
		},
	}
}
