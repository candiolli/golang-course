package main

import "fmt"

func main() {
	/* Aritmeticos */
	soma := 1 + 2
	subtr := 1 - 2
	divisao := 10 / 4
	multipl := 10 * 5
	restoDiv := 10 % 2
	fmt.Println(soma, subtr, divisao, multipl, restoDiv)

	var num1 int16 = 10
	var num2 int16 = 25
	soma2 := num1 + num2
	fmt.Println(soma2)

	/* Atribuição */

	/* Relacionais */
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)

	/* Lógicos */
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	/* Unários */
	numero := 10
	numero++
	numero += 15
	numero--

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	/* Ternarios não existe */

}
