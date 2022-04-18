package main

import "fmt"

// 接口是go语言实现多态的方式
type Dog struct {
	name string
}

func (dog *Dog) eat() {
	fmt.Printf("%s is eating a bone\n", dog.name)
}

type Cat struct {
	name string
}

func (cat *Cat) eat() {
	fmt.Printf("%s is eating a fish\n", cat.name)
}
func main() {
	var dog = Dog{name: "AH"}
	dog.eat()
	var cat = Cat{name: "MIAO"}
	cat.eat()
}
