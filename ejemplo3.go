package main

import "fmt"

func main() {
	var n int

	fmt.Print("Ingrese un número entero positivo: ")
	fmt.Scan(&n)

	original := n
	digitos := 0
	suma := 0

	for n > 0 {
		digito := n % 10
		suma = suma + digito
		digitos = digitos + 1
		n = n / 10
	}

	fmt.Println("Cantidad de dígitos:", digitos)
	fmt.Println("Suma de sus dígitos:", suma)

	switch {
	case digitos == 1:
		fmt.Println("Número de una cifra")
	case digitos == 2:
		fmt.Println("Número de dos cifras")
	case digitos == 3:
		fmt.Println("Número de tres cifras")
	case digitos > 3:
		fmt.Println("Número de varias cifras")
	default:
		fmt.Println("El número debe ser positivo")
	}

	_ = original
}
