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
	{
		fmt.Println("==>", psql.CreateUser)
	}

	port, err := env.GetKey(env.PORT)
	if err != nil {
		log.Fatal(err)
		return
	}

	handler := middlewares.AllowedMethods([]string{"POST"}, signUpHandler{})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}

type signUpHandler struct {
}

func (s signUpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

	id := uuid.New()

	responses.Ok(w, response.SignUp{
		Id:       id.String(),
		Username: reqData.Username,
	})
}
