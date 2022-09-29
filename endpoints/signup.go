package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"

	"example.com/p-service/env"
	"example.com/p-service/middlewares"
	"example.com/p-service/models/request"
	"example.com/p-service/models/response"
	psql "example.com/p-service/prepared_sql"
	"example.com/p-service/responses"
)

const (
	BODY_MAX_BYTES = 100
)

func main() {
	port, err := env.Get(env.PORT)
	if err != nil {
		log.Fatal(err)
		return
	}

	handler := signUp{}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler.handler()))
}

type signUp struct {
}

func (s signUp) handler() http.Handler {
	return middlewares.AllowedMethods([]string{"POST"}, s)
}

func (s signUp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, BODY_MAX_BYTES))
	if err != nil {
		responses.InternalServerError(w)
		return
	}

	var reqData request.SignUp
	err = reqData.UnmarshalBinary(body)
	if err != nil {
		responses.InternalServerError(w)
		return
	}

	fmt.Println("==>", psql.CreateUser)

	id := uuid.New()

	responses.Ok(w, response.SignUp{
		Id:       id.String(),
		Username: reqData.Username,
	})
}
