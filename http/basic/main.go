package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	get()
}
func get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
