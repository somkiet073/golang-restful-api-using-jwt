package demoapi

import (
	"fmt"
	"net/http"
)

func Demo1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Demo 1 API")
}

func Demo2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Demo 2 API")
}
