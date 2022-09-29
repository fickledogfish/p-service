package services

import "database/sql"

type DbService interface {
	Connection() *sql.DB
}

func IsConnected(service DbService) bool {
	return service.Connection().Ping() == nil
}
