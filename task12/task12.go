package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	// Разрешаем запросы с указанных доменов
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Обрабатываем предварительную проверку CORS запросов
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var data struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Hello, %s!", data.Name)
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/api/hello", handlerHello)

	log.Println("Starting server on port 80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
