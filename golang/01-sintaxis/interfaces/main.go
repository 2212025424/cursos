package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet () {
	fmt.Printf("Hola soy %s \n", p.Name)
}

func (p Person) Bye () {
	fmt.Printf("Adios soy %s \n", p.Name)
}

type Text string

func (t Text) Greet () {
	fmt.Printf("Hola soy %s \n", t)
}

func (t Text) Bye () {
	fmt.Printf("Adios soy %s \n", t)
}

func All (gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

func main () {

	p  := Person{Name: "Jose Enrique"}
	var t Text = "un texto"

	All(p, t)
}