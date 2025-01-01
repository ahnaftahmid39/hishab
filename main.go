package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello bro this is awesome")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", rootHandler)

	http.ListenAndServe(":5000", router)
}
