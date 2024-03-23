package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Kei-K23/thunder/pkg/thunder"
)

// var serverPort = 8080

// type dataStruct struct {
// 	Message string `json:"message"`
// }

type resData struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type postData struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {

	start := time.Now() // Record the start time

	// data := postData{
	// 	Title:  "My title again testing",
	// 	Body:   "My body again",
	// 	UserId: 1,
	// }

	resCh, errCh := thunder.HTTPClient("https://jsonplaceholder.typicode.com/posts/1", thunder.Config{
		Method: "GETT",
		Headers: map[string]string{
			"Accept": "application/json",
		},
	})

	res1 := <-resCh
	err := <-errCh

	if err != nil {
		fmt.Println(err)
	}

	if res1 != nil {

		defer res1.Body.Close()
		var res1Data any
		if err := json.NewDecoder(res1.Body).Decode(&res1Data); err != nil {
			panic(err)
		}
		fmt.Println(res1Data)
		// for _, v := range res1Data {
		// 	fmt.Printf("UserId: %d\n", v.PostId)
		// 	fmt.Printf("Id: %d\n", v.Id)
		// 	fmt.Printf("Title: %s\n", v.Name)
		// 	fmt.Printf("Completed: %s\n", v.Email)
		// }
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
