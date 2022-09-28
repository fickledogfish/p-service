package env

import (
	"errors"
	"fmt"
	"os"
)

type key string

const (
	PORT key = "EXPOSED_PORT"
)

func GetKey(key key) (string, error) {
	data, isSet := os.LookupEnv(string(key))
	if !isSet {
		return "", errors.New(fmt.Sprintf("Missing key %s", key))
	}

	return data, nil
}
