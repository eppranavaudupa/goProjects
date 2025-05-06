package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "This method is not supported ", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/form.html")
		return
	}
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err %v", err)
			return
		}
	}
	fmt.Fprintln(w, "Post Request Successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	log.Println(name, address)
	fmt.Fprintf(w, "name is : %s", name)
	fmt.Fprintf(w, "address is : %s", address)
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", hellohandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
