package select_case

import "fmt"

func main()  {
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
		case i1 = <-c1:
			fmt.Printf("recebido ", i1, " de c1 \n")
		case c2 <- i2:
			fmt.Printf("enviado ", i2, " para c2 \n")
		case i3, ok := (<-c3):
			if ok {
				fmt.Printf("recebido ", i3, " de c3 \n")
			} else {
				fmt.Printf("c3 está fechado \n")
			}
		default:
			fmt.Printf(" no comunication \n")
	}
}