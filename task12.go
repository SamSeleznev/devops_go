package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Name struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.ListenAndServe(":80", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var name Name
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, %s!", name.Name)
}
