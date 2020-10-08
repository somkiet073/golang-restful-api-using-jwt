package main

import (
	"golang-restful-api-using-jwt/api/accountapi"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/account/generatekey", accountapi.CreateToken).Methods("POST")
	r.HandleFunc("/api/account/checktoken", accountapi.CheckToken).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err.Error())
	}
}
