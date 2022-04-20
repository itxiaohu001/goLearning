// main.go
package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"os"
)

type data struct {
	Name     string
	Version  string
	Children []*data
}

func newData(n int) *data {
	if n < 0 {
		n = 0
	}
	d := &data{
		Name:     "vuln",
		Version:  fmt.Sprint(rand.Intn(10000)),
		Children: make([]*data, n),
	}
	for i := range d.Children {
		d.Children[i] = newData(n - rand.Intn(n) - 1)
	}

	return d
}

func main() {
	tempTest(newData(10))
}
func tempTest(testdata *data) {
	// 定义一个模板函数
	tmpl, err := template.ParseFiles("./html_use/template/demo1/test.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	f, err := os.Create("test.html")
	if err != nil {
		fmt.Println("create file failed, err", err)
	}
	err = tmpl.Execute(f, testdata)
	defer func() { _ = f.Close() }()
}
