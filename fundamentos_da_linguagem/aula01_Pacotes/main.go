package main

import (
	"CursoGo/fundamentos_da_linguagem/aula01_Pacotes/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(error)

}
