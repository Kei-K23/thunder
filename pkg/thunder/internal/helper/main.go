package helper

import (
	"fmt"
)

// buildURLWithParams constructs the URL with query parameters
func BuildURLWithParams(url string, params map[string]string) string {
	if len(params) == 0 {
		return url
	}

	query := url + "?"
	for k, v := range params {
		query += fmt.Sprintf("%s=%s&", k, v)
	}
	return query[:len(query)-1] // Remove the trailing '&' character
}
