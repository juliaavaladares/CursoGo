package main

import (
	"CursoGo/aplicacao_linha_comando_/app"
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
