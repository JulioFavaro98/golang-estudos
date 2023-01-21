package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipo INTEIRO que pode ser INT8, INT16, INT32 ou INT64
	var numero1 int8 = 127
	fmt.Println(numero1)

	//tipo REAL que pode ser FLOAT32 ou FLOAT64
	var numero2 float32 = 127.45
	fmt.Println(numero2)

	//tipo STRING, no GoLang não existe tipo CHAR
	var variavel string = "Xablau"
	fmt.Println(variavel)

	//tipo BOOLEANO = True or False, se deixar sem resultado a variável será automaticamente FALSE (pois o valor zero no booleano indica condição falsa)
	var booleano1 bool = true
	fmt.Println(booleano1)
	//False
	var booleano2 bool
	fmt.Println(booleano2)

	//Tipo ERRO se declara ERROR e tem que instanciar a função errors.New para poder atribuir um valor ao erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
