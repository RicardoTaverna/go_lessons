package main

import "fmt"

func main() {
	a := 100
	b := 200
	var ret int

	ret = max(a, b)

	fmt.Printf("O valor máximo é: %d \n", ret)

}

// função para retornar o maior valor entre dois numeros
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
