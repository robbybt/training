package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

//func for FiveLastRequestDetail
func SetArraylistBelowFive(arr *[]map[string]interface{}, val map[string]interface{}) {
	array := *arr
	if len(array) == 5 {
		array = append(array[1:len(array)], val)
		return
	}
	array = append(array, val)
	*arr = array
}

func PrintOut(data interface{}) {
	t, _ := json.MarshalIndent(data, "", "\t")
	createFile(t)
}

func createFile(data []byte) {
	t := time.Now().Format("2006-01-02 15:04:05")
	a, _ := os.Create(fmt.Sprint(t, ".txt"))
	a.Write(data)
}
