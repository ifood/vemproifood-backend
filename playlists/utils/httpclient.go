package utils

import (
	"net/http"
)

// HTTPClient : used to define how a http client implementation should behave
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// HTTPClientDo : type to implement a Do function to the http client
type HTTPClientDo func(*http.Request) (*http.Response, error)
