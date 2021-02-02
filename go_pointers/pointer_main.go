package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("Endereço de memória da variavel 'a': %x \n", &a)

	b := 20
	var ip *int

	ip = &b
	fmt.Printf("Endereço de memória da variavel 'b': %x \n", &b)

	fmt.Printf("Endereço salvo em uma variável do tipo ponteiro: %x \n", ip)

	// Acessando a variável usando um ponteiro
	fmt.Printf("Valor da variável *ip: %d \n", *ip)

	var ptr *int

	if ptr == nil {
		fmt.Println("O ponteiro *prt é nulo")
	}

}
