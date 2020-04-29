package utils

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// HTTPClient : used to define how a http client implementation should behave
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// HTTPClientDo : type to implement a Do function to the http client
type HTTPClientDo func(*http.Request) (*http.Response, error)

type httpClientMock struct {
	mock.Mock
	funcDo HTTPClientDo
}

// NewHTTPClientMock : create a generic http client for mocks
func NewHTTPClientMock(do HTTPClientDo) HTTPClient {
	return &httpClientMock{
		funcDo: do,
	}
}

// Do : used to make requests
func (c *httpClientMock) Do(r *http.Request) (*http.Response, error) {
	return c.funcDo(r)
}
