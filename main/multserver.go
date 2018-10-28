package main

import (
	"net/http"
	"fmt"
)
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTP Server1")
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTP Server2")
}
func main()  {
	go func(){
		mux := http.NewServeMux()
		mux.HandleFunc("/", handler2)
		http.ListenAndServe(":1234", mux)
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler1)
	http.ListenAndServe(":12345", mux)
}
