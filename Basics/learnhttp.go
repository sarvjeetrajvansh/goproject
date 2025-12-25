package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
