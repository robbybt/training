package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

//func for FiveLastRequestDetail
func SetFiveLastRequestDetail(arr *[]map[string]interface{}, val map[string]interface{}) {
	array := *arr
	if len(array) == 5 {
		array = append(array[1:len(array)], val)
		return
	}
	array = append(array, val)
	*arr = array
}

//PrintOut Create file with json indent
func PrintOut(data interface{}) {
	t, _ := json.MarshalIndent(data, "", "\t")
	createFile(t)
}

// Create file
func createFile(data []byte) {
	t := time.Now().Format("2006-01-02 15:04:05")
	a, _ := os.Create(fmt.Sprint(t, ".txt"))
	a.Write(data)
}

//GetSinceTimeStart get time lapse from start endpoint
func GetSinceTimeStart(ctx context.Context) time.Duration {
	val := ctx.Value(KeyTimeStart).(time.Time)
	return time.Since(val)
}

func RetryFunc(f func() error) {
	go func() {
		for i := 0; i < 5; i++ {
			<-time.Tick(time.Second * time.Duration(i+1))
			err := f()
			if err == nil {
				break
			}
		}
		fmt.Println("retry finish")
	}()
}
