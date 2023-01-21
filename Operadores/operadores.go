package main

import "fmt"

func main() {
	//Aritméticos
	soma := 5 + 10
	sub := 15 - 3
	mult := 3 * 10
	div := 15 / 3
	resto := 15 % 2
	fmt.Println(soma, sub, mult, div, resto)

	//Atribuição
	var txt1 string = "String"
	txt2 := "String2"
	fmt.Println(txt1, txt2)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Lógicos
	fmt.Println("_________________________________________")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //Operador AND
	fmt.Println(verdadeiro || falso) //Operador OR
	fmt.Println(!verdadeiro)         //Operador NEGAÇÃO
	fmt.Println(!falso)

	//Unários
	fmt.Println("_________________________________________")
	numero := 10
	numero++ // numero + 1
	fmt.Println(numero)
	numero += 15 // numero + 15
	fmt.Println(numero)

	numero-- // numero - 1
	fmt.Println(numero)
	numero -= 15 // numero - 15
	fmt.Println(numero)

	numero *= 3 // numero * 3
	fmt.Println(numero)
	numero /= 3 // numero / 3
	fmt.Println(numero)
	numero %= 3 // numero % 3
	fmt.Println(numero)
}
