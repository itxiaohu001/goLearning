package main

import "fmt"

// 接口是go语言实现多态的方式
type animal interface {
	eat()
}
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
func whichAnimal(x animal) {
	x.eat()
}
func main() {
	dog := &Dog{name: "ah"}
	cat := &Cat{name: "miao"}
	whichAnimal(dog)
	whichAnimal(cat)
}
