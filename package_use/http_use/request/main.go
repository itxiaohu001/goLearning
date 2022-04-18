package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 如何设置查询参数参数
	requestByParams()
	//如何定制请求头
	requestByHead()
}
func printBody(resp *http.Response) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
func requestByParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	// here
	params := make(url.Values)
	params.Add("name", "hyf")
	params.Add("add", "18")
	request.URL.RawQuery = params.Encode()
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
	printBody(r)
}

func requestByHead() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	// here
	request.Header.Add("user-agent", "chrome")
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
