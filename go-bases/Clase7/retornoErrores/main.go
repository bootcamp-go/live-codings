package main

import "fmt"

type HTTPError struct {
	Code        int
	Description string
}

func (h *HTTPError) Error() string {
	return fmt.Sprintf("status code %d: %s", h.Code, h.Description)
}

func request() error {
	return &HTTPError{404, "not found"}
}
func main() {
	err := request()
	if err != nil {
		// an error occured
		if err.(*HTTPError).Code == 404 {
			// handle a "not found" error
		} else {
			// handle a different error
		}
	}
}
