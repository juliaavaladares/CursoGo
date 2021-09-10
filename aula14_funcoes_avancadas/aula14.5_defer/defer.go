package main

import "fmt"

func funcao1() {
	fmt.Println(" Executou funcao 1")
}

func funcao2() {
	fmt.Println(" Executou funcao 2")
}


func alunoEstaAprovado(nota1, nota2 float32) bool  {
	defer fmt.Println("Média calculada. Resultado retornado")
	media := (nota1 + nota2)/2

	if media >= 6{
		return true
	}
	return false
}


func main(){

	defer funcao1() //DEFER = ADIAR
	funcao2()

	fmt.Println(alunoEstaAprovado(4,8))

}