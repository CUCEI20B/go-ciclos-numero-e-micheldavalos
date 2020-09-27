package main

import "fmt"

func factorial(numero float64) float64 {
	if numero == 0 {
		return 1
	}
	return numero * factorial(numero-1)
}

func main() {
	var x float64
	for {
		// fmt.Print("Indica el nímero límite para calcular el número E: ")
		fmt.Scan(&x)
		if x >= 0 {
			break
		}
	}

	var i float64
	var resultado float64
	for i = 0; i <= x; i++ {
		resultado += 1 / factorial(i)
	}

	fmt.Println(resultado)
}
