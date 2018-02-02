package main

import (
	"fmt"
	"net/http"

	"github.com/go-examples/http_sample/handler"
)

func main() {
	err := handler.Setup()
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/user", handler.UserHandler)
	http.ListenAndServe(":8080", nil)
}
