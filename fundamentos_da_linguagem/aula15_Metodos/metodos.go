package main

import "fmt"

type usuario struct{
	nome string
	idade int
}

func (u *usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}


func (u usuario) verificaIdade() bool{
	return u.idade >=18
}


func (u *usuario) fazerAniversario() {
	u.idade++
}

func main()  {
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Julia", 22}
	usuario2.salvar()

	maiorDeIdade := usuario2.verificaIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}