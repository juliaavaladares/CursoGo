package main

import "fmt"

func recuperarExecucao()  {
	if r := recover(); r != nil{
		fmt.Println("Funcao recuperada com sucesso")
	}
}

func alunoEstaAprovado(nota1, nota2 float32) bool  {
	defer recuperarExecucao()
	media := (nota1 + nota2)/2

	if media > 6{
		return true
	} else if media < 6 {
		return false
	}
	panic("A MÉDIA É EXATAMENTE 6!")
}


func main(){

	fmt.Println(alunoEstaAprovado(6,6))
	fmt.Println("Pos Execucao")


}