package main

import (
	"fmt"
	"strings"
)

func main() {
	saludo("Enrique", "Zempo")

	emoji := "carita"
	change(&emoji)
	fmt.Printf("\nEs un: %s", emoji)

	total := suma(3, 5)
	fmt.Printf("\nEl resultado es: %d", total)

	texto := "JosE EnRiQue"
	tmin, tmay := convert(texto)
	fmt.Println("\n", tmin, tmay)

	nums := []int{1, 4, 6, 23, 5, 78, 20, 3, 56, 5, 3, 37}
	result := filter(nums, func(num int) bool {
		return num <= 10
	})
	fmt.Println(result)

	saludo := hello("Enrique")
	fmt.Println(saludo())

	fmt.Println("El resultado de la suma es: ", sum(3, 5, 10, 4, 2))

	// Funcion anonima (funciones que no poseen nombre)
	anonima := func(text string) {
		fmt.Println(text)
	}
	anonima("Soy una funcion anonima")

	// Funcion anonima autoejecutada (funciones que no poseen nombre)
	func() {
		fmt.Println("Funcion anonima autojecutada")
	}()
}

// Funcion variatica (parametros no definidos)
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// Funcion que retorna funcion
func hello(name string) func() string {
	return func() string {
		return "hello " + name
	}
}

// Funcion que reciben funciones como parametro
func filter(nums []int, callback func(int) bool) []int {
	result := []int{}

	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	return result
}

// Funcion con multiples retornos
func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}

// Funcion con retorno
func suma(num1, num2 int) int {
	return num1 + num2
}

// Funcion con parametros por referencia
func change(value *string) {
	*value = "animal"
}

//Funcion con parametros por valor
func saludo(fr_name string, ls_name string) {
	fmt.Printf("Hola %s %s... !", fr_name, ls_name)
}
