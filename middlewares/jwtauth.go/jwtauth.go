package jwtauth

import (
	"golang-restful-api-using-jwt/api/accountapi"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "MySecret"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("key")
		if tokenString == "" {
			accountapi.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
		} else {
			tokenString := r.Header.Get("key")
			result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err == nil && result.Valid {
				next.ServeHTTP(w, r)
			} else {
				accountapi.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
			}
		}
	})
}
