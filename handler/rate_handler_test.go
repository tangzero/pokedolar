package handler_test

import (
	"io/ioutil"
	"net/http/httptest"
	"pokedolar/handler"
	"testing"
)

func TestGetRate(t *testing.T) {
	request := httptest.NewRequest("GET", "/rate?from=USD&to=BRL", nil)
	response := httptest.NewRecorder()

	handler.GetRate(response, request)

	result := response.Result()
	if result.StatusCode != 200 {
		t.Errorf("expected 200. got %d", result.StatusCode)
	}

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	content := string(bytes)

	if content != "5.44" {
		t.Errorf("expected 5.44. got %s", content)
	}
}