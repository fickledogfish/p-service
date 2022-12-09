package network

import (
	"fmt"
	"log"
	"net/http"

	"example.com/p-service/env"
)

func RunHandler(handler http.Handler) {
	port, err := env.Get(env.PORT)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}
