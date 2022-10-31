package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Http)
	http.ListenAndServe(":8000", nil)
}

func Http(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
