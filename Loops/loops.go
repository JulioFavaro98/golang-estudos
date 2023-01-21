package main

import (
	"fmt"
	"time"
)

func main() {
	//Primeira maneira de utilizar Loop
	i := 0
	for i < 10 {
		fmt.Println("Incrementando I", i+1)
		i++
		time.Sleep(time.Second)
	}

	//Segunda maneira de utilizar Loop
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J", j+1)
		time.Sleep(time.Second)
	}

	//Terceira maneira de utilizar Loop
	nomes := [3]string{"Leo do Lau", "Blau Blau", "Ematerius"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "COCAINE" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"Nome":      "Leonardo",
		"Sobrenome": "Salvalaggio",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
