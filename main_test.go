package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestGetRateEndpoint(t *testing.T) {
	defaultServer := http.DefaultServeMux
	server := httptest.NewServer(defaultServer)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	expect.GET("/rate").
		WithQuery("from", "USD").
		WithQuery("to", "BRL").
		Expect().
		Status(http.StatusOK).
		Text().Equal("5.44")
}
