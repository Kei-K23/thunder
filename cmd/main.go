package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var serverPort = 8080

type dataStruct struct {
	Message string `json:"message"`
}

func main() {
	go startServer()

	time.Sleep(100 * time.Millisecond)

	requestData := []byte(`{"client_message": "hello from client"}`)
	jsonRequestData := bytes.NewReader(requestData)

	requestURL := fmt.Sprintf("http://localhost:%d/?id=1234", serverPort)

	req, err := http.NewRequest(http.MethodPost, requestURL, jsonRequestData)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer sdjfljfalfj")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var body dataStruct
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		panic(err)
	}

	fmt.Println("client: got response")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	fmt.Printf("client: response data: %s\n", body.Message)
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("server: %s\n", r.Method)

		fmt.Printf("server: %s /\n", r.Method)
		fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
		fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
		fmt.Printf("server: headers:\n")
		for headerName, headerValue := range r.Header {
			fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
		}

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("server: could not read request body: %s\n", err)
		}
		fmt.Printf("server: request body: %s\n", reqBody)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataStruct{Message: "hello!"})
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("error running server: %s\n", err)
	}
}
