package main

import "fmt"

func main() {
	// Formas de declarar variáveis
	var var1 string = "var 1"
	var2 := 22
	fmt.Println(var1, var2)

	var (
		var3 string = "tete"
		var4 string = "222"
	)

	fmt.Println(var3, var4)

	var5, var6 := "ddd", 333

	fmt.Println(var5, var6)

	// Constantes
	const const1 string = "teste const"
	fmt.Println(const1)

	// Troca de valores entre variaveis
	var3, var4 = var4, var3
	fmt.Println(var3, var4)
}
