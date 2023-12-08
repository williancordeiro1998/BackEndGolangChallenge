package bitcoin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateBalance(t *testing.T) {
	// Caso de teste 1
	wallet := Wallet{Address: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", Balance: 100.0}
	amount := 10.0
	newBalance := wallet.CalculateBalance(amount)
	assert.Equal(t, newBalance, 90.0, "Saldo incorreto após uma transação de 10.0")

	// Caso de teste 2
	wallet.Balance = 50.0
	amount = 60.0
	newBalance = wallet.CalculateBalance(amount)
	assert.LessOrEqual(t, newBalance, 0.0, "O saldo não deve ser negativo")

}

func TestIsValidTransaction(t *testing.T) {
	// Caso de teste 1
	sender := Wallet{Address: "SenderAddress", Balance: 100.0}
	recipient := Wallet{Address: "RecipientAddress", Balance: 50.0}
	transaction := Transaction{Sender: sender, Recipient: recipient, Amount: 30.0}
	isValid := IsValidTransaction(transaction)
	assert.True(t, isValid, "Transação válida foi considerada inválida")

	// Caso de teste 2
	sender.Balance = 20.0
	transaction = Transaction{Sender: sender, Recipient: recipient, Amount: 30.0}
	isValid = IsValidTransaction(transaction)
	assert.False(t, isValid, "Transação inválida foi considerada válida")

}

// Estruturas fictícias para ilustrar os testes
type Wallet struct {
	Address string
	Balance float64
}

func (w *Wallet) CalculateBalance(amount float64) float64 {
	return w.Balance - amount
}

type Transaction struct {
	Sender    Wallet
	Recipient Wallet
	Amount    float64
}

func IsValidTransaction(transaction Transaction) bool {
	return transaction.Sender.Balance >= transaction.Amount
}
