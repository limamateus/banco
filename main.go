package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroDaConta int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	poderSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if poderSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito efetuado com sucesso"
	} else {
		return "Falha ao depositar, verifique o valor informado!"
	}

}

func main() {

	contaDoMateus := ContaCorrente{}
	contaDoMateus.titular = "Mateus"
	contaDoMateus.numeroDaConta = 123
	contaDoMateus.numeroAgencia = 12344
	contaDoMateus.saldo = 500.0

	fmt.Println(contaDoMateus.saldo)
	fmt.Println(contaDoMateus.Sacar(200))
	fmt.Println(contaDoMateus.saldo)
	fmt.Println(contaDoMateus.Depositar(2000))
	fmt.Println(contaDoMateus.saldo)
	fmt.Println(contaDoMateus.Sacar(1000))
	fmt.Println(contaDoMateus.saldo)

}
