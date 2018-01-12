package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 1")
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":80", nil)
}
