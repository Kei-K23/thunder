package thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Kei-K23/thunder/pkg/thunder/internal/helper"
)

// Config holds configuration options for HTTP requests
type Config struct {
	Method      string // HTTP method like
	Params      map[string]string
	Headers     map[string]string
	JSONPayload interface{}       // JSON payload data
	FormPayload map[string]string // Form payload data
}

// HTTPClient makes an HTTP request with the provided configuration
func HTTPClient(url string, config Config) (chan *http.Response, chan error) {
	resCh := make(chan *http.Response)
	errCh := make(chan error)

	go func() {
		// Build URL with query parameters if provided
		reqURL := helper.BuildURLWithParams(url, config.Params)
		fmt.Println(reqURL)

		// Create request based on the specified method
		var req *http.Request
		var err error

		switch config.Method {
		case http.MethodGet:
			req, err = http.NewRequest(http.MethodGet, reqURL, nil)
		case http.MethodPost:
			req, err = buildPostRequest(reqURL, config)
		default:
			err = fmt.Errorf("unsupported HTTP method: %s", config.Method)
		}

		if err != nil {
			resCh <- nil
			errCh <- err // Send the error to the error channel
			return
		}

		// Set request headers
		for k, v := range config.Headers {
			req.Header.Set(k, v)
		}

		// Create HTTP client with timeout
		client := http.Client{
			Timeout: 30 * time.Second,
		}

		// Send request and handle response
		res, err := client.Do(req)
		if err != nil {
			resCh <- nil
			errCh <- err // Send the error to the error channel
			return
		}

		resCh <- res
	}()

	return resCh, errCh
}

// buildPostRequest builds a POST request with the specified payload type
func buildPostRequest(url string, config Config) (*http.Request, error) {
	var payloadData []byte
	var err error

	switch {
	case config.JSONPayload != nil:
		payloadData, err = json.Marshal(config.JSONPayload)
	case len(config.FormPayload) > 0:
		payloadData = []byte(formEncode(config.FormPayload))
	default:
		return nil, fmt.Errorf("no payload data provided")
	}

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json") // Set default content type for POST requests

	return req, nil
}

// formEncode encodes form payload data into URL encoded format
func formEncode(data map[string]string) string {
	var encodedData string
	for key, value := range data {
		encodedData += fmt.Sprintf("%s=%s&", key, value)
	}
	return encodedData[:len(encodedData)-1] // Remove the trailing '&'
}
