package main

import (
	"log"
	"net/http"
)

const ADDR = ":8080"

func main() {
	handler := signUpHandler{}

	log.Fatal(http.ListenAndServe(ADDR, handler))
}

type signUpHandler struct {
}

func (s signUpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
}
