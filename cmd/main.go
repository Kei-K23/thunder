package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {

	data := postData{
		Title: "Patch request body",
	}

	resCh, errCh := thunder.HTTPClient("https://jsonplaceholder.typicode.com/postss/1", thunder.Config{
		Method: http.MethodPatch,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
		JSONPayload: data,
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
		if res1.StatusCode != 200 {
			fmt.Println("error")
			fmt.Println(res1Data)
		}
	}

}
