package main

import (
	"html/template"
	"log"
	"os"
)

type Recipient struct {
	Name, Gift string
	Attended   bool
}

// Define a template.
const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.{{else}}
It is a shame you couldn't make it to the wedding.{{end}}
{{with .Gift}}Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

func main() {
	var recSlice = []Recipient{
		{"xiaohu", "airpods", true},
		{"xiaohui", "", false},
		{"xiaozhou", "iphone", false},
	}
	t := template.Must(template.New("letter").Parse(letter))
	for _, r := range recSlice {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println(err)
		}
	}
}
