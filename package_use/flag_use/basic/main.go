package main

import (
	"flag"
	"fmt"
)

var (
	Name string
	Age  int
)

func init() {
	flag.StringVar(&Name, "name", "DefaultName", "Input your name")
	flag.IntVar(&Age, "age", 0, "Input your age")
}
func main() {
	flag.Parse()
	fmt.Println(Name)
	fmt.Println(Age)
}
