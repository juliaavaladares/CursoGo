package main

import "fmt"

func main()  {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome": "Julia",
		"sobrenome": "Valadares",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro": "Julia",
			"segundo": "Valadares",
		}, 
		"curso":{
			"nome": "Ciencia da Computacao",
			"campus": "Juiz de Fora",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Peixes",
	}

	fmt.Println(usuario2)
}