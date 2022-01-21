package main

import (
	"app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start server")

	aplicacao := app.Gerar()

	err := aplicacao.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
