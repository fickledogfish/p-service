package middlewares

import (
	"net/http"

	"example.com/p-service/responses"
)

func AllowedMethods(methods []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			responses.Forbidden(w, "Forbidden method")
			return
		}

		next.ServeHTTP(w, req)
	})
}
