package main

import "fmt"

func main() {
	var capitalPaisMap map[string]string
	capitalPaisMap = make(map[string]string)

	capitalPaisMap["França"] = "Paris"
	capitalPaisMap["Italia"] = "Roma"
	capitalPaisMap["Japão"] = "Tokyo"
	capitalPaisMap["India"] = "Nova Delhi"

	for pais := range capitalPaisMap {
		fmt.Println("Capital da", pais, "é", capitalPaisMap[pais])
	}

	capital, ok := capitalPaisMap["Estados Unidos"]
	if ok {
		fmt.Println("Capital dos Estados Unidos é", capital)
	} else {
		fmt.Println("Capital dos Estados Unidos não está presente")
	}

	delete(capitalPaisMap, "França")
	fmt.Println()
	fmt.Println("Registro da França foi deletado")
	fmt.Println()
	fmt.Println("map Atualizado")
	fmt.Println()

	for pais := range capitalPaisMap {
		fmt.Println("Capital da", pais, "é", capitalPaisMap[pais])
	}

}
