package main

import (
	"fmt"
	"main/clientes"
	"main/contas"
)

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteMateus := clientes.Titular{Nome: "Mateus", CPF: "46063928863", Profissao: "QA"}

	contaPoupanca := contas.ContaPoupaca{Titular: clienteMateus, NumeroAgencia: 123, NumeroDaConta: 1320}

	contaPoupanca.Depositar(300)

	fmt.Println("Saldo é : ", contaPoupanca.ObterSaldo())

	contaPoupanca.Sacar(100)

	fmt.Println("Saldo é : ", contaPoupanca.ObterSaldo())

	fmt.Println(contaPoupanca)

	PagarBoleto(&contaPoupanca, 60)

	fmt.Println(contaPoupanca)

}
