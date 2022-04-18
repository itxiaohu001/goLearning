package main

import (
	"fmt"
	"html"
)

func main() {
	res := html.EscapeString(`"Fran & Freddie's Diner" <tasty@example.com>`)
	fmt.Println(res)
	res = html.UnescapeString(res)
	fmt.Println(res)
}
