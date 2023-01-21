package main

import "fmt"

func main() {
	//Declarando variável com o tipo de dado explicito
	var variavel1 string = "Variavel 1"
	//Declarando variável sem o tipo de dado
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Declarando mais de uma variável utilizando a tipagem de cada uma
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	//Declarando mais de uma variável sem utilizar a tipagem
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//Invertendo o valor das variáveis sem utilizar uma variável auxiliar
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
