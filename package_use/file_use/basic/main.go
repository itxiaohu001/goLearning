package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	n, err := io.WriteString(f, "hello,go!")
	fmt.Println(n, err)
}
