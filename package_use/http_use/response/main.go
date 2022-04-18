package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	//responseBody(r)
	//status(r)
	//header(r)
	encoding(r)
}
func responseBody(r *http.Response) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
func status(r *http.Response) {
	fmt.Println(r.StatusCode)
	fmt.Println(r.Status)
}
func header(r *http.Response) {
	fmt.Println(r.Header.Get("Content-Type"))
}
func encoding(r *http.Response) {
	// 1.content-type中会提供编码信息，比如 content-type="text/html_use;charset=utf-8"
	// 2.html_use head meta获取编码,
	// <meta http-equiv=Content-Type content="text/html_use;charset=utf-8">
	// 3.以上都没有编码信息的话就用猜测的方式来取到编码信息
	bufReader := bufio.NewReader(r.Body)
	bytes, err := bufReader.Peek(1024)
	if err != nil {
		panic(err)
	}
	// DetermineEncoding通过检查确定HTML文档的编码
	e, _, _ := charset.DetermineEncoding(bytes, r.Header.Get("content-type"))
	// 调用编码对应的解码器解码
	bodyReader := transform.NewReader(bufReader, e.NewDecoder())
	content, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
