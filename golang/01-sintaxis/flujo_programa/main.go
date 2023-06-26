package main

import "fmt"

/* DEFER posterga las instrucciones hasta antes de retornar la funcion */
/* DEFER guarda las istrucciones bajo una estructura LIFO */
/* Con DEFER se guada el estado inicial de los parametros */

func main() {
	defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)
}
