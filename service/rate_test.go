package service_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pokedolar/service"
	"testing"
)


func TestRatings(t *testing.T) {
	t.Cleanup(func() {
		http.DefaultClient.Transport = nil
	})

	cases := []testCase{
		{from: "USD", to: "BRL", expected: 5.39},
		{from: "USD", to: "EUR", expected: 0.76},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("from_%s_to_%s", c.from, c.to), func(t *testing.T) {
			http.DefaultClient.Transport = &mockHTTP{c}

			rate, err := service.Rate(c.from, c.to)

			if err != nil {
				t.Fatal("failed to get the rate.", err)
			}

			if rate != c.expected {
				t.Fatalf("invalid rating. expected %f. got %f", c.expected, rate)
			}
		})
	}
}

type testCase struct{
	from string
	to string
	expected float64
}

type mockHTTP struct {
	test testCase
}

func (m *mockHTTP) RoundTrip(r *http.Request) (*http.Response, error) {
	query := r.URL.Query()
	if query.Get("base") == m.test.from && query.Get("symbols") == m.test.to {
		rec := httptest.NewRecorder()
		rec.Code = 200
		fmt.Fprintf(rec.Body, "{ \"rates\": { \"%s\": %f } }", m.test.to, m.test.expected)
		return rec.Result(), nil
	}
	return nil, errors.New("not found")
}
