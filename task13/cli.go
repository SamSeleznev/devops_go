package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: cli <server_address> <name>")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
	name := os.Args[2]

	// Создаем запрос
	requestData := map[string]string{"name": name}
	requestBody, _ := json.Marshal(requestData)
	requestURL := fmt.Sprintf("http://%s/api/hello", serverAddress)
	request, _ := http.NewRequest("GET", requestURL, bytes.NewBuffer(requestBody))

	// Выполняем запрос
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error sending GET request: ", err)
	}
	defer response.Body.Close()

	// Обрабатываем ответ
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	fmt.Println(string(responseBody))
}
