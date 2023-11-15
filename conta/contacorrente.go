package conta

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroDaConta int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	poderSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if poderSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Deposito efetuado com sucesso", c.Saldo
	} else {
		return "Falha ao depositar, verifique o valor informado!", c.Saldo
	}

}

func (c *ContaCorrente) Tranferir(valorDatranferencia float64, contaDistino *ContaCorrente) bool {
	if c.Saldo >= valorDatranferencia && c.Saldo < 0 {
		c.Saldo -= valorDatranferencia
		contaDistino.Depositar(valorDatranferencia)
		return true
	} else {
		return false

	}
}
