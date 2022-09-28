package main

import (
	"log"
	"net/http"

	"example.com/p-service/responses"
)

const ADDR = ":8080"

func main() {
	handler := signUpHandler{}

	log.Fatal(http.ListenAndServe(ADDR, handler))
}

type signUpHandler struct {
}

func (s signUpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		responses.Forbidden(w, "Forbidden method")
	}
}
