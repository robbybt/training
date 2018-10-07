package main

import (
	"net/http"
)

//HitCount each 10 sec
var HitCount int64

//ErrorCount each 10 sec
var ErrorCount int64

//FiveLastRequestDetail will save Five LastRequestDetail
var FiveLastRequestDetail []map[string]interface{}

func init() {
	FiveLastRequestDetail = make([]map[string]interface{}, 0)
}

func main() {
	http.Handle("/training", middleware(GetTrainingHandler().Handler))
	http.ListenAndServe("localhost:9090", nil)
}

func middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}

func setzz(a *[]int64) {
	*a = append(*a, 10)
}
