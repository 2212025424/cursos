package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {}

func (p Paypal) Pay () {
	fmt.Println("Pagado desde paypal")
}

type Cash struct {}

func (c Cash) Pay () {
	fmt.Println("Pagado desde efectivo")
}

type CreditCard struct {}

func (c CreditCard) Pay () {
	fmt.Println("Pagado desde tarjeta de credito")
}

func Factory (method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default: 
		return nil
	}
} 

func main () {

	var method uint

	fmt.Println("Ingrese un numero de pago: ")
	_, err := fmt.Scanln(&method)

	if err != nil {
		panic("debe digitar un metodo válido")
	} 
	
	if method > 3 {
		panic("debe ingresar un valor de las opciones")
	}

	paymethod := Factory(method)
	paymethod.Pay()
}