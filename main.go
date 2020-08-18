package main

import (
	"net/http"
	"pokedolar/handler"
	"pokedolar/service"
)

func init() {
	http.HandleFunc("/rate", handler.GetRate(service.Rate))
}

func main() {
	http.ListenAndServe(":3000", nil)
}