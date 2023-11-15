package contas

import "main/clientes"

type ContaPoupaca struct {
	Titular                                clientes.Titular
	NumeroAgencia, NumeroDaConta, Operacao int
	saldo                                  float64
}

func (c *ContaPoupaca) Sacar(valorDoSaque float64) string {
	poderSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if poderSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupaca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito efetuado com sucesso", c.saldo
	} else {
		return "Falha ao depositar, verifique o valor informado!", c.saldo
	}

}

func (c *ContaPoupaca) Tranferir(valorDatranferencia float64, contaDistino *ContaPoupaca) bool {
	if c.saldo >= valorDatranferencia && c.saldo < 0 {
		c.saldo -= valorDatranferencia
		contaDistino.Depositar(valorDatranferencia)
		return true
	} else {
		return false

	}
}

func (c *ContaPoupaca) ObterSaldo() float64 {
	return c.saldo
}
