package response

import "encoding/json"

type errorResponse struct {
	Message string `json:"message,omitempty"`
}

func NewError(message string) errorResponse {
	return errorResponse{
		Message: message,
	}
}

func (e errorResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}
