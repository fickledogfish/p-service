package response

import (
	"encoding/json"
)

type SignUp struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (s SignUp) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
