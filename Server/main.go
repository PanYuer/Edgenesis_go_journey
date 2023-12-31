// This is an http server. When you visit "/hello", it prints "Hello world! Welcome to Edgenesis!"
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! Welcome to Edgenesis!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
