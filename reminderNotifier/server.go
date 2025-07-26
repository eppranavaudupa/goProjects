package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO WORLD", w)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
