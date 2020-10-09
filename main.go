package main

import (
	"golang-restful-api-using-jwt/api/accountapi"
	"golang-restful-api-using-jwt/api/demoapi"
	"golang-restful-api-using-jwt/middlewares/jwtauth.go"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/account/generatekey", accountapi.CreateToken).Methods("POST")
	r.HandleFunc("/api/account/checktoken", accountapi.CheckToken).Methods("GET")

	r.Handle("/api/demo/demo1", jwtauth.JWTAuth(http.HandlerFunc(demoapi.Demo1))).Methods("GET")
	r.Handle("/api/demo/demo2", jwtauth.JWTAuth(http.HandlerFunc(demoapi.Demo2))).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err.Error())
	}
}
