package main

import (
	"html/template"
	"os"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{Material: "wool", Count: 17}
	tmpl, err := template.New("test").Parse(`{{.Count}} of {{.Material}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
