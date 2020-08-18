package service_test

import (
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"net/http"
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
			defer gock.Off() // Flush pending mocks after test execution

			gock.New("https://api.ratesapi.io").
				Get("/api/latest").
				MatchParam("base", c.from).
				MatchParam("symbols", c.to).
				Reply(200).
				BodyString(fmt.Sprintf("{ \"rates\": { \"%s\": %f } }", c.to, c.expected))

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
