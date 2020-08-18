package service

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func Rate(from string, to string) (float64, error) {
	url := fmt.Sprintf("https://api.ratesapi.io/api/latest?base=%s&symbols=%s", from, to)

	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	var r struct {
		Rates map[string]float64
	}

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return 0, err
	}

	rate := r.Rates[to]
	return math.Trunc(rate * 100) / 100, nil
}