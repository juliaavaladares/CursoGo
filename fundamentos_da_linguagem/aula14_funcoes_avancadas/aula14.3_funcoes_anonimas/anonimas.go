package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá Mundo!")
	}()

	//func (texto string)  {
	//	fmt.Println(texto)
	//}

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
