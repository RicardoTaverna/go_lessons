package main

import "fmt"

func main()  {
	grade := "G"
	nota := 30

	switch nota {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50, 60, 70: grade = "C"
		default: grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("Excelente! \n")
	
	case grade == "B", grade == "C":
		fmt.Printf("Muito bom! \n")
	
	case grade == "D":
		fmt.Printf("Você passou! \n")
	
	case grade == "F":
		fmt.Printf("Melhor tentar novamente! \n")
	
	default:
		fmt.Printf("Grade inválida! \n")
		
	}

	// utilizando o tipo da variavel
	var x interface{}
     
	switch i := x.(type) {
	   case nil:	  
		  fmt.Printf("tipo de x é:%T \n", i)                
	   case int:	  
		  fmt.Printf("x é int \n")                       
	   case float64:
		  fmt.Printf("x é float64 \n")           
	   case func(int) float64:
		  fmt.Printf("x é func(int) \n")                      
	   case bool, string:
		  fmt.Printf("x é bool ou string \n")       
	   default:
		  fmt.Printf("Não faço ideia que type seja \n")     
	}

}