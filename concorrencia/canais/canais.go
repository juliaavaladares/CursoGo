package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Ola mundo", canal)

	for {
		mensagem, aberto := <-canal // recebendo um valor, esperando que o canal receba um valor
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}
