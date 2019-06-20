package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/something", handleSomething)
	fmt.Println("starting server..")
	http.ListenAndServe("0.0.0.0:9090", router)
}

func handleSomething(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("hello world"))
}