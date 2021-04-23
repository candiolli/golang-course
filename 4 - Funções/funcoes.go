package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calcMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtr := n1 - n2
	return soma, subtr
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		return "meu nome é " + txt
	}

	fmt.Println(f("silas"))

	resultSoma, resultSubtr := calcMatematicos(10, 20)
	fmt.Println(resultSoma, resultSubtr)

	resultSoma2, _ := calcMatematicos(10, 20)
	fmt.Println(resultSoma2)
}
