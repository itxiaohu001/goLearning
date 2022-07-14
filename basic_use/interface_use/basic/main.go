package main

import "fmt"

// 接口是go语言实现多态的方式
type animal interface {
	eat()
}
type Dog struct {
}
type Cat struct {
}

func (dog *Dog) eat() {
	fmt.Printf("dog is eating")
}
func (cat *Cat) eat() {
	fmt.Println("cat is eating")
}
func newAnimal(str string) animal {
	if str == "dog" {
		return &Dog{}
	} else if str == "cat" {
		return &Cat{}
	} else {
		return nil
	}
}
func main() {
	a := newAnimal("dog")
	a.eat()
}
