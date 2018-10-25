package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//HitCount each 10 sec
type HitCountType struct {
	count int64
	mutex sync.Mutex
}

var HitCount HitCountType

//ErrorCount each 10 sec
type ErrorCountType struct {
	count int64
	mutex sync.Mutex
}

var ErrorCount ErrorCountType

//FiveLastRequestDetail will save Five LastRequestDetail
type FiveLastRequestDetailType struct {
	detail []map[string]interface{}
	mutex  sync.Mutex
}

var FiveLastRequestDetail FiveLastRequestDetailType

func init() {
	FiveLastRequestDetail.detail = make([]map[string]interface{}, 0)
	ClearCountBackgroundProccess()
	CheckErrorReportingBackgroundProccess()
}

type MyHandler func(ctx context.Context, w http.ResponseWriter, r *http.Request) (resp interface{}, err error)

func main() {
	http.Handle("/training", middleware(GetTrainingHandler().Handler))
	http.ListenAndServe("localhost:9090", nil)
}

func middleware(h MyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HitCount.Set(HitCount.Get() + 1)
		ctx := r.Context()
		ctx = context.WithValue(ctx, KeyTimeStart, time.Now())
		resp, err := h(ctx, w, r)
		if err != nil {
			if resp != nil {
				FiveLastRequestDetail.SetWithFunc(func() {
					SetFiveLastRequestDetail(&FiveLastRequestDetail.detail, resp.(map[string]interface{}))
				})
			}
			ErrorCount.Set(ErrorCount.Get() + 1)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprint(HitCount.Get(), " hit count ", resp.(string))))
	}
}

func (h *HitCountType) Set(v int64) {
	h.mutex.Lock()
	h.count = v
	h.mutex.Unlock()
}

func (h *HitCountType) Get() int64 {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	return h.count
}

func (e *ErrorCountType) Set(v int64) {
	e.mutex.Lock()
	e.count = v
	e.mutex.Unlock()
}

func (e ErrorCountType) Get() int64 {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.count
}

func (f *FiveLastRequestDetailType) SetWithFunc(funcSet func()) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	funcSet()
}

func (f *FiveLastRequestDetailType) Get() []map[string]interface{} {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return f.detail
}

func ClearCountBackgroundProccess() {
	go func() {
		for {
			<-time.Tick(time.Second * 100)
			HitCount.Set(0)
			ErrorCount.Set(0)
		}
	}()
}

func CheckErrorReportingBackgroundProccess() {
	go func() {
		for {
			<-time.Tick(time.Second)
			if ErrorCount.Get() > 5 {
				PrintOut(FiveLastRequestDetail.Get())
				HitCount.Set(0)
				ErrorCount.Set(0)
				<-time.Tick(time.Second * 5)
			}
		}
	}()
}
