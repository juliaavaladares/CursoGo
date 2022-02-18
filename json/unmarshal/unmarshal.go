package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	cachorroEmJSon := `{"nome":"rex","raca":"dalmata","idade":3}`
	var cachorro Cachorro

	fmt.Println("Cachorro json: ", cachorroEmJSon)
	fmt.Println("Cachorro vazio: ", cachorro)

	err := json.Unmarshal([]byte(cachorroEmJSon), &cachorroEmJSon)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Cachorro preenchido: ", cachorro)

}
