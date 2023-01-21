package main

import "fmt"

//Criando uma struct
type usuario struct {
	//Propriedades da struct
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo structs")
	//Instanciando a struc
	var u usuario
	u.nome = "Xampson"
	u.idade = 22
	fmt.Println(u)

	//Segundo modo de instanciar a struct com inferencia de tipo
	usuario2 := usuario{"Xablau", 24}
	fmt.Println(usuario2)

	//Instanciando com apenas uma das propriedades da struct
	usuario3 := usuario{nome: "Maconheiro"}
	fmt.Println(usuario3)
	usuario4 := usuario{idade: 24}
	fmt.Println(usuario4)

}
