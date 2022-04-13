package main

import "fmt"

type Animal interface {
	eat()
}

type AnimalFactory interface {
	create(name string) Animal
}

type Tiger struct {
	name string
}

func (t *Tiger) eat() {
	fmt.Println(t.name, " eat...")
}

type TigerFactory struct {
}

func (f *TigerFactory) create(name string) Animal {
	return &Tiger{name: name}
}

func main() {
	var f AnimalFactory = &TigerFactory{}
	var a = f.create("tiger")
	a.eat()
}
