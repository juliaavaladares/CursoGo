package main

import "fmt"

type user struct{
	nome string
	idade uint
	endereco adress
}

type adress struct{
	logadouro string
	numero uint
}

func main()  {
	fmt.Println("Arquivo Structs")

	var usuario user
	fmt.Println(usuario)
	usuario.idade = 22
	usuario.nome = "Julia"
	fmt.Println(usuario)

	endereco := adress{"Rua X", 0}

	secondUser := user{"Julia", 22, endereco}
	fmt.Println(secondUser)

	thirdUser := user{nome: "Julia"}
	fmt.Println(thirdUser)
}