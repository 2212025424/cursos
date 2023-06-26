package main

import "fmt"

type Person struct {
	Name string
	Age uint
}

func NewPerson (name string, age uint) Person {
	return Person{name, age}
}

func (p *Person) Greet () {
	fmt.Println("Hello am a Person ! ")
}

type Human struct {
	Age uint
	Clidren uint
}

func NewHuman (age, clidren uint) Human {
	return Human{age, clidren}
}

type Employee struct {
	Person
	Human
	Salary float64
}

func (e *Employee) Greet () {
	fmt.Println("Hello am a Employee")
}

func NewEmployee (name string, age uint, clidren uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, clidren), salary}
}

func (e *Employee) Payroll () {
	fmt.Println(e.Salary * 0.90)
}

func main () {
	e := NewEmployee("Alejandro", 30, 2, 1000000)
	fmt.Println(e.Name, e.Person.Age)
	e.Greet()
	e.Payroll()
	fmt.Println("Que onda como estas !!!")
}