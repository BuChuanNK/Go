package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int 'json:"page"'
	Fruits []string 'json:"fruits"'
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
}