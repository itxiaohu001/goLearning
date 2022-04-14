package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	requestByParams()
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

	params := make(url.Values)
	params.Add("name", "hyf")
	params.Add("add", "18")

	request.URL.RawQuery = params.Encode()

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	printBody(r)
}
