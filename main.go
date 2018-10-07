package main

import (
	"context"
	"net/http"
	"time"
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

type MyHandler func(ctx context.Context, w http.ResponseWriter, r *http.Request) (resp interface{}, err error)

func main() {
	http.Handle("/training", middleware(GetTrainingHandler().Handler))
	http.ListenAndServe("localhost:9090", nil)
}

func middleware(h MyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, KeyTimeStart, time.Now())
		resp, err := h(ctx, w, r)
		if err != nil {
			SetFiveLastRequestDetail(&FiveLastRequestDetail, resp.(map[string]interface{}))
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(resp.(string)))
	}
}

func setzz(a *[]int64) {
	*a = append(*a, 10)
}
