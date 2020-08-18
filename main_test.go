package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRateEndpoint(t *testing.T) {
	defaultServer := http.DefaultServeMux
	server := httptest.NewServer(defaultServer)
	defer server.Close()

	response, err := server.Client().Get(server.URL + "/rate?from=USD&to=BRL")
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 200 {
		t.Errorf("expected 200. got %d", response.StatusCode)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	content := string(bytes)

	if content != "5.44" {
		t.Errorf("expected 5.44. got %s", content)
	}
}