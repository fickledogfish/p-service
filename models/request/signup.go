package request

import (
	"bytes"
	"encoding/json"
)

type SignUp struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email,omitempty"`
}

func (s *SignUp) UnmarshalBinary(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	return decoder.Decode(&s)
}
