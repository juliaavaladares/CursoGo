package main

import (
	"fmt"
)

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func calculos(n1, n2 int)  (int, int){
	soma := n1+n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main()  {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func (txt string)  {
		fmt.Println(txt)
	}

	f("Texto 1")

	resultadoSoma, resultadoSubtracao := calculos(10,15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}