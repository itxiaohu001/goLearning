package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.OpenFile(`./test.txt`, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer f.Close()
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	// 设置光标，偏移量为5，相对位置为文件开头
	// 也就是说从文件头开始的第五个位置后开始读
	fmt.Println("seek=>")
	_, err = f.Seek(5, 0)
	if err != nil {
		fmt.Println("set seek err", err)
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read file err", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("truncate=>")
	err = f.Truncate(0)
	if err != nil {
		fmt.Println("truncate err", err)
		return
	}
	_, err = f.Write([]byte{32, 32})
	if err != nil {
		fmt.Println("write err", err)
		return
	}
	b, err = ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read file err", err)
		return
	}
	fmt.Println(string(b))
}
