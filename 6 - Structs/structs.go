package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
	bairro     string
	cidade     string
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{"Jairo", 21, endereco{"rua 15", 1, "B", "Sao leo"}}
	fmt.Println(u2)

	u3 := usuario{idade: 22}
	fmt.Println(u3)
}
