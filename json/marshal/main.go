package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	cachorro := Cachorro{
		Nome:  "rex",
		Raca:  "dalmata",
		Idade: 3,
	}

	fmt.Println(cachorro)

	cachorroJson, err := json.Marshal(cachorro)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Cachoorro em JSON: ", cachorroJson)
	fmt.Println(bytes.NewBuffer(cachorroJson))
}
