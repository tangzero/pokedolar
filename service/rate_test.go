package service_test

import (
	"fmt"
	"pokedolar/service"
	"testing"
)


func TestRatings(t *testing.T) {
	cases := []struct{
		from string
		to string
		expected float64
	}{
		{from: "USD", to: "BRL", expected: 5.44},
		{from: "USD", to: "EUR", expected: 0.83},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("from_%s_to_%s", c.from, c.to), func(t *testing.T) {
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


