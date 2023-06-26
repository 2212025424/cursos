package main

import "fmt"

type course struct {
	name string
}

func (c course) Print () {
	fmt.Printf("%+v\n", c)
}

// declaracion de alias
type myAlias = course

// definicion de tipo
type newCourse course

type newBool bool 

func (n newBool) test () string {
	if (n) {
		return "VERDADERO"
	}
	return "FALSO"
}

func main () {
	var variable newBool = false
	fmt.Println(variable.test())
	
	dato := course{"nombre del curso"}
	dato.Print()
	fmt.Println("Que onda !")
}