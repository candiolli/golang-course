package main

import (
	"errors"
	"fmt"
)

func main() {
	/* Valores texto e char(retorna valor do caracter na tabela ASC) */
	str2 := "texto"
	char := 'B'
	fmt.Println(str2, char)

	// Valores inteiros
	var num int = 1000000000000000000
	var num16 int16 = 10000
	var num32 int32 = 1000000000
	var num64 int64 = 1000000000000000000
	fmt.Println(num, num16, num32, num64)

	/* Alias para tipos
	   INT32 = RUNE
	   UINT8 = BYTE */
	var num3 rune = 123456
	var num4 byte = 123
	fmt.Println(num3, num4)

	/* Valores monetarios */
	numReal1 := 123.25
	var numReal2 float64 = 22222.22
	fmt.Println(numReal1, numReal2)

	/* Valor Zero */
	var numero int
	var texto string
	var bool1 bool
	fmt.Println(numero, texto, bool1)

	/* Boolean */
	var bool2 bool = true
	fmt.Println(bool2)

	/* Erro */
	var erro error
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro, erro2)
}
