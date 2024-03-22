package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Kei-K23/thunder/pkg/thunder"
)

// var serverPort = 8080

// type dataStruct struct {
// 	Message string `json:"message"`
// }

type resData struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// func main() {

// 	start := time.Now() // Record the start time

// 	res1, err := thunder.Get("https://jsonplaceholder.typicode.com/todos/1", thunder.Config{
// 		Headers: map[string]string{
// 			"Content-Type": "application/json",
// 			"Accept":       "application/json",
// 		},
// 	})

// 	if err != nil {
// 		defer res1.Body.Close()
// 		var res1Data resData
// 		if err := json.NewDecoder(res1.Body).Decode(&res1Data); err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Response 1:")
// 		fmt.Printf("UserId: %d\n", res1Data.UserId)
// 		fmt.Printf("Id: %d\n", res1Data.Id)
// 		fmt.Printf("Title: %s\n", res1Data.Title)
// 		fmt.Printf("Completed: %t\n", res1Data.Completed)
// 	}

// 	res2, err := thunder.Get("https://jsonplaceholder.typicode.com/todos/2", thunder.Config{
// 		Headers: map[string]string{
// 			"Content-Type": "application/json",
// 			"Accept":       "application/json",
// 		},
// 	})

// 	if err != nil {
// 		defer res2.Body.Close()
// 		var res2Data resData
// 		if err := json.NewDecoder(res2.Body).Decode(&res2Data); err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Response 2:")
// 		fmt.Printf("UserId: %d\n", res2Data.UserId)
// 		fmt.Printf("Id: %d\n", res2Data.Id)
// 		fmt.Printf("Title: %s\n", res2Data.Title)
// 		fmt.Printf("Completed: %t\n", res2Data.Completed)
// 	}

// 	fmt.Println("hello fetch")

// 	elapsed := time.Since(start) // Calculate elapsed time
// 	fmt.Printf("Total time taken: %s\n", elapsed)
// }

func main() {
	ch1 := make(chan *http.Response)

	start := time.Now() // Record the start time

	go thunder.Get("https://jsonplaceholder.typicode.com/todos/1", thunder.Config{
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	}, ch1)

	res1 := <-ch1

	if res1 != nil {
		defer res1.Body.Close()
		var res1Data resData
		if err := json.NewDecoder(res1.Body).Decode(&res1Data); err != nil {
			panic(err)
		}
		fmt.Println("Response 1:")
		fmt.Printf("UserId: %d\n", res1Data.UserId)
		fmt.Printf("Id: %d\n", res1Data.Id)
		fmt.Printf("Title: %s\n", res1Data.Title)
		fmt.Printf("Completed: %t\n", res1Data.Completed)
	}
	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Total time taken: %s\n", elapsed)
}

// func startServer() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("server: %s\n", r.Method)

// 		fmt.Printf("server: %s /\n", r.Method)
// 		fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
// 		fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
// 		fmt.Printf("server: headers:\n")
// 		for headerName, headerValue := range r.Header {
// 			fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
// 		}

// 		reqBody, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Printf("server: could not read request body: %s\n", err)
// 		}
// 		fmt.Printf("server: request body: %s\n", reqBody)

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(dataStruct{Message: "hello!"})
// 	})
// 	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("server: %s\n", r.Method)

// 		fmt.Printf("server: %s\n", r.Method)
// 		fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
// 		fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
// 		fmt.Printf("server: headers:\n")
// 		for headerName, headerValue := range r.Header {
// 			fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
// 		}

// 		reqBody, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Printf("server: could not read request body: %s\n", err)
// 		}
// 		fmt.Printf("server: request body: %s\n", reqBody)

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(dataStruct{Message: "hello!"})
// 	})

// 	server := http.Server{
// 		Addr:    fmt.Sprintf(":%d", serverPort),
// 		Handler: mux,
// 	}
// 	if err := server.ListenAndServe(); err != nil {
// 		fmt.Printf("error running server: %s\n", err)
// 	}
// }
