package main

import "fmt"

func main() {
	var n int

	fmt.Print("Ingrese un número entero positivo: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}

	if n%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}

	switch {
	case n >= 1 && n <= 5:
		fmt.Println("Número pequeño")
	case n >= 6 && n <= 10:
		fmt.Println("Número mediano")
	case n > 10:
		fmt.Println("Número grande")
	default:
		fmt.Println("El número debe ser positivo")
	}
}
