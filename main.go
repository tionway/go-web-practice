package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	fmt.Println("Hello World!")
	server.ListenAndServe()
}
