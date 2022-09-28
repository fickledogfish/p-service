package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"

	"example.com/p-service/models/request"
	"example.com/p-service/models/response"
	"example.com/p-service/responses"
)

const (
	ADDR = ":8080"

	BODY_MAX_BYTES = 100
)

func main() {
	handler := signUpHandler{}

	log.Fatal(http.ListenAndServe(ADDR, handler))
}

type signUpHandler struct {
}

func (s signUpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		responses.Forbidden(w, "Forbidden method")
		return
	}

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
