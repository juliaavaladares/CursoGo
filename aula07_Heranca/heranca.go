package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade uint
	altura uint
}

type estudante struct{
	pessoa 
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa1 := pessoa{"Julia", "Valadares",22, 150}
	fmt.Println(pessoa1)
	estudante1 := estudante{pessoa1, "Ciencia da Computacao", "UFJF"}
	fmt.Println(estudante1)
}