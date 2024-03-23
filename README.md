# Thunder ⚡️

Thunder is a Go package for making HTTP requests with various payload types and configurations.

## Features

- Supports GET, POST, PUT, PATCH, and DELETE methods ✨.
- Supports JSON, Form URL-encoded, and Multipart/form-data payloads ✨.
- Customizable headers and URL parameters ✨.

## Installation

```bash
go get github.com/Kei-K23/thunder
```

## Usage

### Basic HTTP GET request

```bash
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/thunder"
)

func main() {
	// Example of making a GET request
	resCh, errCh := thunder.HTTPClient("https://api.example.com", thunder.Config{
		Method:  http.MethodGet,
		Headers: map[string]string{"Authorization": "Bearer token"},
	})

	res := <-resCh
	err := <-errCh
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response Status:", res.Status)
}
```

## Examples

### Making a POST request with JSON payload

```bash
resCh, errCh := thunder.HTTPClient("https://api.example.com", thunder.Config{
	Method:     http.MethodPost,
	Headers:    map[string]string{"Content-Type": "application/json"},
	JSONPayload: map[string]interface{}{"key": "value"},
})

```

### Making a POST request with Form URL-encoded payload

```bash
resCh, errCh := thunder.HTTPClient("https://api.example.com", thunder.Config{
	Method:     http.MethodPost,
	Headers:    map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
	FormPayload: map[string]string{"key": "value"},
})

```

### Making a POST request with Multipart/form-data payload

```bash
resCh, errCh := thunder.HTTPClient("https://api.example.com", thunder.Config{
	Method:            http.MethodPost,
	MultipartPayload:  map[string]string{"key": "value"},
})

```

### Making a PUT request with JSON payload

```bash
resCh, errCh := thunder.HTTPClient("https://api.example.com/resource/123", thunder.Config{
	Method:      http.MethodPut,
	Headers:     map[string]string{"Content-Type": "application/json"},
	JSONPayload: map[string]interface{}{"key": "new_value"},
})

```

### Making a PATCH request with JSON payload

```bash
resCh, errCh := thunder.HTTPClient("https://api.example.com/resource/123", thunder.Config{
	Method:      http.MethodPatch,
	Headers:     map[string]string{"Content-Type": "application/json"},
	JSONPayload: map[string]interface{}{"key": "updated_value"},
})

```

### Making a DELETE request

```bash
resCh, errCh := thunder.HTTPClient("https://api.example.com/resource/123", thunder.Config{
	Method:      http.MethodDelete,
	Headers:     map[string]string{"Authorization": "Bearer token"},
})

```

## Contribution

Thank you for considering contributing to the Thunder package! Contributions are welcomed and appreciated. Please read guide for contribution [CONTRIBUTION](./CONTRIBUTION.md).

## License

This project is licensed under the [MIT License](./LICENSE) - see the LICENSE file for details.
