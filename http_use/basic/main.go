package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	get()
	post()
	put()
	del()
}

// Here are a few simple demonstrations
func get() {
	// http://httpbin.org A website that demonstrates http_use
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func post() {
	if r, err := http.Post("http://httpbin.org/post", "", nil); err != nil {
		panic(err)
	} else {
		defer func() {
			err := r.Body.Close()
			if err != nil {
				panic(err)
			}
		}()
		if content, err := ioutil.ReadAll(r.Body); err != nil {
			panic(err)
		} else {
			fmt.Println(string(content))
		}
	}
}

func put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
func del() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
