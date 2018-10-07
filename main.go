package main

import (
	"net/http"
)

func main() {
	http.Handle("/training", middleware(GetTrainingHandler().Handler))
	http.ListenAndServe("localhost:9090", nil)
}

func middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
