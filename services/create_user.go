package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"

	"example.com/p-service/env"
	dtos "example.com/p-service/services/DTOs"
)

type NewUserService interface {
	Create(ctx context.Context, user dtos.CreateNewUser) (dtos.CreatedUser, error)
}

func NewNewUserService() (NewUserService, error) {
	service := newUserService{}
	if err := service.getEnv(); err != nil {
		return service, err
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		service.host,
		service.port,
		service.user,
		service.password,
		service.dbname,
	)

	connector, err := pq.NewConnector(connString)
	if err != nil {
		return service, err
	}

	service.dbConnection = sql.OpenDB(connector)
	return service, err
}

type newUserService struct {
	timeout time.Duration

	dbConnection *sql.DB

	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (service *newUserService) getEnv() (err error) {
	service.host, err = env.Get(env.PGDB_HOST)
	if err != nil {
		return
	}

	service.port, err = env.Get(env.PGDB_PORT)
	if err != nil {
		return
	}

	service.user, err = env.Get(env.PGDB_USER)
	if err != nil {
		return
	}

	service.password, err = env.Get(env.PGDB_PASSWORD)
	if err != nil {
		return
	}

	service.dbname, err = env.Get(env.PGDB_DBNAME)
	if err != nil {
		return
	}

	return nil
}

func (s newUserService) Create(
	ctx context.Context,
	user dtos.CreateNewUser,
) (dtos.CreatedUser, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return dtos.CreatedUser{}, nil
}

func (s newUserService) Connection() *sql.DB {
	return s.dbConnection
}
