package main

import "fmt"

type Livros struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Livro1 Livros
	var Livro2 Livros

	// especificação do livro 1
	Livro1.title = "Go Programming"
	Livro1.author = "Mahesh Kumar"
	Livro1.subject = "Go Programming Tutorial"
	Livro1.book_id = 654987

	// especificações do livro 2
	Livro2.title = "Python Programming"
	Livro2.author = "Zara Ali"
	Livro2.subject = "Python for dummies"
	Livro2.book_id = 789654

	printBook(Livro1)
	printBook(Livro2)
}

func printBook(livro Livros) {
	fmt.Println("===================================")
	fmt.Printf("ID: %d \n", livro.book_id)
	fmt.Printf("Titulo: %s \n", livro.title)
	fmt.Printf("Autor: %s \n", livro.author)
	fmt.Printf("Assunto: %s \n", livro.subject)
	fmt.Println("===================================")
}
