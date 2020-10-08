package accountapi

import (
	"encoding/json"
	"fmt"
	"golang-restful-api-using-jwt/database/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "MySecret"

func CreateToken(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	// Binding value to account
	err := json.NewDecoder(r.Body).Decode(&account)
	// check error
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		fmt.Printf("err accountapi line 22: %s", err.Error())
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": account.Username,
			"password": account.Password,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		if tokenString, err2 := token.SignedString([]byte(secretKey)); err2 != nil {
			ResponseWithError(w, http.StatusBadRequest, err2.Error())
			fmt.Printf("err accountapi line 32: %s", err2.Error())
		} else {
			ResponseWithJson(w, http.StatusOK, model.JWTToken{Token: tokenString})
		}

	}
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("key")
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil && result.Valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("InValid")
	}
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	ResponseWithJson(w, code, map[string]string{"error": msg})
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
