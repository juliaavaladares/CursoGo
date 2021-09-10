package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1{
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci (posicao - 1)
}

func main(){
	
	fmt.Println("Funções Recursivas")

	posicao := 12

	for i := 1; i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}