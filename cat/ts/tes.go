package main

import "fmt"

type Namer interface {
	Name() string
}

type Animal struct {
	name string
}

func (a Animal) Name() string {
	return "Животное: " + a.name
}

type Student struct {
	name string
	age  int
}

func (s Student) Name() string {
	return s.name
}

func (s Student) Age() int {
	return s.age
}

func hello(n Namer) {
	fmt.Printf("Привет, %s!\n", n.Name())
}

func main() {
	s := Student{name: "Полина"}
	a := Animal{name: "Котик"}

	hello(s)
	hello(a)
}
