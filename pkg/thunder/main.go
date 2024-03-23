package thunder

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kei-K23/thunder/pkg/thunder/internal/helper"
)

// requestData := []byte(`{"client_message": "hello from client"}`)
// jsonRequestData := bytes.NewReader(requestData)

// requestURL := fmt.Sprintf("http://localhost:%d/?id=1234", serverPort)

// req, err := http.NewRequest(http.MethodPost, requestURL, jsonRequestData)
// if err != nil {
// 	panic(err)
// }

// req.Header.Set("Content-Type", "application/json")
// req.Header.Set("Authorization", "Bearer sdjfljfalfj")

// client := &http.Client{
// 	Timeout: 30 * time.Second,
// }

// res, err := client.Do(req)

// if err != nil {
// 	panic(err)
// }

// defer res.Body.Close()

// var body dataStruct
// if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
// 	panic(err)
// }

// fmt.Println("client: got response")
// fmt.Printf("client: status code: %d\n", res.StatusCode)
// fmt.Printf("client: response data: %s\n", body.Message)

// Config holds configuration options for HTTP requests
type Config struct {
	Params  map[string]string
	Headers map[string]string
}

// Get makes a GET request with the provided URL and configuration
func Get(url string, config Config) chan *http.Response {
	// init channel to receive response or error if not successful
	resCh := make(chan *http.Response)

	go func() {
		reqUrl := helper.BuildURLWithParams(url, config.Params)

		fmt.Println(reqUrl)
		req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
		if err != nil {
			resCh <- nil
			return
		}

		for k, v := range config.Headers {
			req.Header.Set(k, v)
		}

		client := http.Client{
			Timeout: 30 * time.Second,
		}

		res, err := client.Do(req)
		if err != nil {
			resCh <- nil
			return
		}
		resCh <- res
	}()
	return resCh
}
