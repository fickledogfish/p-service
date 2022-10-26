package middlewares

import (
	"net/http"

	"example.com/p-service/responses"
	"example.com/p-service/utils"
)

func AllowedMethods(allowedMethods []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !utils.Contains(allowedMethods, req.Method) {
			responses.Forbidden(w, "Forbidden method")
			return
		}

		next.ServeHTTP(w, req)
	})
}
