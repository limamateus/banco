package main

import (
	"banco/conta"
	"fmt"
)

func main() {

	contaDoMautes := conta.ContaCorrente{Titular: "Mateus"}

	fmt.Println(contaDoMautes)

}
