package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"Nome":      "Majin",
		"Sobrenome": "Boo",
	}
	fmt.Println(usuario["Nome"])
	fmt.Println(usuario["Sobrenome"])
}
