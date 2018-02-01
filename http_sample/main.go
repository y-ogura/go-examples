package main

import (
	"net/http"

	"github.com/go-examples/http_sample/handler"
)

func main() {
	http.HandleFunc("/user", handler.UserHandler)
	http.ListenAndServe(":8080", nil)
}
