package thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Config holds configuration options for HTTP requests
type Config struct {
	Method           string            // HTTP method
	Params           map[string]string // URL parameters
	Headers          map[string]string // Request headers
	JSONPayload      interface{}       // JSON payload data
	FormPayload      map[string]string // Form payload data
	MultipartPayload map[string]string // Multipart/form-data payload data
}

// HTTPClient makes an HTTP request with the provided configuration
func HTTPClient(url string, config Config) (chan *http.Response, chan error) {
	// Initialize channels for response and error
	resCh := make(chan *http.Response)
	errCh := make(chan error)

	// Start go routine to make concurrent requests
	go func() {
		// Create request based on the specified method
		var req *http.Request
		var err error

		// Build URL with query parameters if provided
		reqURL := buildURLWithParams(url, config.Params)

		// Check request method and perform HTTP request
		switch config.Method {
		case http.MethodGet:
			req, err = http.NewRequest(config.Method, reqURL, nil)
		case http.MethodPost:
			req, err = buildRequest(reqURL, config)
		case http.MethodPut:
			req, err = buildRequest(reqURL, config)
		case http.MethodPatch:
			req, err = buildRequest(reqURL, config)
		case http.MethodDelete:
			if config.FormPayload != nil {
				req, err = buildRequest(reqURL, config)
			} else if config.JSONPayload != nil {
				req, err = buildRequest(reqURL, config)
			} else if config.MultipartPayload != nil {
				req, err = buildRequest(reqURL, config)
			} else {
				req, err = http.NewRequest(config.Method, url, nil)
			}
		default:
			err = fmt.Errorf("unsupported HTTP method: %s", config.Method)
		}

		if err != nil {
			resCh <- nil
			errCh <- err // Send the error to the error channel
			return
		}

		// Set request headers if provide by user
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
		errCh <- err
	}()

	return resCh, errCh
}

// buildRequest builds a POST request with the specified payload type
func buildRequest(reqUrl string, config Config) (*http.Request, error) {
	var req *http.Request
	var err error

	switch {
	case config.JSONPayload != nil:
		// JSON payload
		payloadData, err := json.Marshal(config.JSONPayload)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(config.Method, reqUrl, bytes.NewBuffer(payloadData))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

	case len(config.FormPayload) > 0:
		// Form payload (application/x-www-form-urlencoded)
		formData := url.Values{}
		for key, value := range config.FormPayload {
			formData.Set(key, value)
		}
		req, err = http.NewRequest(config.Method, reqUrl, strings.NewReader(formData.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	case len(config.MultipartPayload) > 0:
		// Multipart/form-data payload
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for key, value := range config.MultipartPayload {
			part, err := writer.CreateFormField(key)
			if err != nil {
				return nil, err
			}
			_, err = part.Write([]byte(value))
			if err != nil {
				return nil, err
			}
		}

		err = writer.Close()
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(config.Method, reqUrl, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

	default:
		return nil, fmt.Errorf("no payload data provided")
	}

	return req, nil
}

func buildURLWithParams(url string, params map[string]string) string {
	if len(params) == 0 {
		return url
	}

	query := url + "?"
	for k, v := range params {
		query += fmt.Sprintf("%s=%s&", k, v)
	}
	return query[:len(query)-1] // Remove the trailing '&' character
}
