package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Starting web server on port 8000")

	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Response from the helloHandler\n")
	io.WriteString(w, "URL.Path = "+r.URL.Path)
}
