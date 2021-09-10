package main

import "fmt"

func closure()  func(){
	texto := "Dentro da função closure"

	funcao := func ()  {
		fmt.Println(texto)
	}
	return funcao
}


func main(){

	texto := "dentro da funçao main"
	fmt.Println(texto)
	fmt.Println("-------------------------")

	funcaoNova := closure()
	funcaoNova()

}