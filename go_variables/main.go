package main

import "fmt"

func main()  {
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x é do tipo %T \n", x)	

	// Declaração de variáveis dinâmica e inferência
	var y float64 = 30.0
	fmt.Println(y)
	fmt.Printf("y é do tipo %T \n", y)

	z := 42
	fmt.Println(z)
	fmt.Printf("z é do tipo %T \n", z)

	// Declarações mixtas
	var a, b, c = 3, 5, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a é do tipo %T \n", a)
	fmt.Printf("b é do tipo %T \n", b)
	fmt.Printf("c é do tipo %T \n", c)

	// Declaração de constantes
	const LENGHT int = 10
	const WIDTH int = 5
	var area int

	area = LENGHT * WIDTH
	fmt.Printf("O valor da área: %d \n", area)

	// Retonar o endereço da memória de uma variável
	fmt.Printf("Endereço memória da variável 'a': %d \n", &a)
}