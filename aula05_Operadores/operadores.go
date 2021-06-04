package main

import "fmt"

func main()  {
	//ARITMETICOS
	soma := 1+2
	subtracao := 1 -2
	divisao := 10/4
	multiplicacao := 10*5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao,divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 700
	
	//soma2 := numero1+numero2 -> ERRO, não é possível fazer operações com tipos diferentes
	soma2 := numero1+ int16(numero2)
	fmt.Println(soma2)

	//FIM ARITMÉTICOS

	//ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)
	//FIM ATRIBUICAO

	//OPERADORES RELACIONAIS
	fmt.Println("-------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	fmt.Println(1 > 2)
	fmt.Println("-------------------")

	//FIM DOS RELACIONAIS

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println("-------------------")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	fmt.Println("-------------------")
	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	fmt.Println("-------------------")
	fmt.Println(numero)

	numero += 15
	fmt.Println(numero)

	numero--
	fmt.Println(numero)
	fmt.Println("-------------------")
	//FIM OPERADORES UNÁRIOS




}