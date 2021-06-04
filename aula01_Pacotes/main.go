package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(error)

}