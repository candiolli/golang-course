package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa := pessoa{"Silas", "Silva", 31, 1.65}
	estudante := estudante{pessoa, "Eng Civil", "Feevale"}

	fmt.Println(estudante)
}
