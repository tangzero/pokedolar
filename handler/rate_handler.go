package handler

import (
	"fmt"
	"net/http"
	"pokedolar/service"
)

func GetRate(w http.ResponseWriter, r *http.Request) {
	rate, err := service.Rate(r.URL.Query().Get("from"), r.URL.Query().Get("to"))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprint(w, rate)
}