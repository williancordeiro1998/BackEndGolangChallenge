// bitcoin_exception.go
package exception

// BitcoinException é uma exceção personalizada para erros relacionados ao Bitcoin
type BitcoinException struct {
	message string
}

// NewBitcoinException cria uma nova instância de BitcoinException com a mensagem fornecida
func NewBitcoinException(message string) *BitcoinException {
	return &BitcoinException{message: message}
}

// Error retorna a mensagem de erro associada à exceção
func (be *BitcoinException) Error() string {
	return be.message
}
