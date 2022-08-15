package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(fmt.Sprintf("Hello Dogus!\n")))
	})
	http.ListenAndServe(":80", nil)
}
