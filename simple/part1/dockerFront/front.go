package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func test(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://back.default:7777")
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(w, "Cannot reach BACK for some reasons!!!\n")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", test)

	http.ListenAndServe("0.0.0.0:6666", nil)
}
