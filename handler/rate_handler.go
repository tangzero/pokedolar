package handler

import (
	"fmt"
	"net/http"
	"pokedolar/service"
)

func GetRate(rateFunc service.RateFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rate, err := rateFunc(r.URL.Query().Get("from"), r.URL.Query().Get("to"))
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
		}
		fmt.Fprint(w, rate)
	}
}