package env

import (
	"errors"
	"fmt"
	"os"
)

type key string

const (
	PORT key = "EXPOSED_PORT"

	PGDB_HOST     key = "POSTGRES_HOST"
	PGDB_PORT     key = "POSTGRES_PORT"
	PGDB_USER     key = "POSTGRES_USER"
	PGDB_PASSWORD key = "POSTGRES_PASSWORD"
	PGDB_DBNAME   key = "POSTGRES_DBNAME"
)

func Get(key key) (string, error) {
	data, isSet := os.LookupEnv(string(key))
	if !isSet {
		return "", errors.New(fmt.Sprintf("Missing key %s", key))
	}

	return data, nil
}
