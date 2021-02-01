package main

import "fmt"

func main() {
	var a int
	b := 15

	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("O valor de 'a' é: %d \n", a)
	}

	for a < b {
		a++
		fmt.Printf("Agora, o valor de 'a' é: %d \n", a)
	}

	for i, x := range numbers {
		fmt.Printf("Valor de x = %d em %d \n", x, i)
	}

}
