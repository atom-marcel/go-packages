package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}

type response1 struct {
	Page    int    `json:"page"`
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, req *http.Request) {

	res := &response1{
		Page:    1,
		Message: "Hello World!",
	}

	resJSON, _ := json.MarshalIndent(res, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
