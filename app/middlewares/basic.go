package middlewares

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
)

func SetAuthBasic(next http.HandlerFunc, username string, password string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Basic Authentication credentials
		u, p, hasAuth := r.BasicAuth()
		isValid := hasAuth && u == username && p == password
		if !isValid {
			err := response_mapper.NewError(response_mapper.ErrUnauthorized, response_mapper.NewResponseMultiLang(
				response_mapper.MultiLanguages{
					ID: "Token tidak valid",
					EN: "Invalid token",
				},
			))

			response_mapper.RenderJSON(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
