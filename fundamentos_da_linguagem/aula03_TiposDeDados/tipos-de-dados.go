package main

import (
	"errors"
	"fmt"
)

func main()  {
	// O GO possui os seguintes tipos de numeros inteiros:
	// int int8, int16, int32, int64
	var numero int64 = 1000000000
	fmt.Println(numero)

	// uint -> unsigned int (número inteiro sem sinal)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//Alias 
	//INT32 = RUNE 
	var numero3 rune = 1234
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMEROS RAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123333333333333333333333333.45
	fmt.Println(numeroReal2)

	//FIM NUMEROS REAIS

	//STRINGS
	var str string = "Texto"
	str2 := "Texto 2"

	fmt.Println(str, str2)

	char := 'A'
	fmt.Println(char)

	//FIM STRINGS

	texto := 5
	fmt.Println(texto)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)






	
	
}