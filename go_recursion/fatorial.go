package main

import "fmt"

func fatorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * fatorial(i-1)
}

func main() {
	i := 15
	fmt.Printf("Fatorial de %d é %d \n", i, fatorial(i))

}
