package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the other side.\n")
}

func main() {
	http.HandleFunc("/", test)
	
	http.ListenAndServe("0.0.0.0:7777", nil)
}
