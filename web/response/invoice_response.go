package response

type PostResponse struct {
	OK int `json:"ok"`
}

func NewPostResponse() PostResponse {
	return PostResponse{
		OK: 1,
	}
}
