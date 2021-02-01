package main

import (
	"fmt"
	"strings"
)

func main() {
	saudacao := "Hello world!"

	fmt.Printf("string normal: %s", saudacao)
	fmt.Println()
	fmt.Printf("Bytes: ")

	for i := 0; i < len(saudacao); i++ {
		fmt.Printf("%x ", saudacao[i])
	}

	fmt.Println()

	const textoExemplo = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Printf("Quoted String: %+q", textoExemplo)
	println()

	fmt.Printf("O comprimento da string 'saudacao' é: %d \n", len(saudacao))

	frase := []string{"Olá", "mundo!"}
	fmt.Println(strings.Join(frase, " "))
}
