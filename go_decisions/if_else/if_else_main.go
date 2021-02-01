package main

import "fmt"

func main()  {
	a := 20
	b := 100

	if(a < 20){
		fmt.Printf("'a' é menor que 20 \n")
	} else if(a == 20){
		fmt.Printf("'a' é igual a 20 \n")
	} else {
		fmt.Printf("'a' é maior que 20 \n")
	}
	fmt.Printf("O valor de a é: %d \n", a)

	if(a == 20){
		if(b == 100){
			fmt.Printf("Ganhou na loteria \n")
		}
	}
	
}