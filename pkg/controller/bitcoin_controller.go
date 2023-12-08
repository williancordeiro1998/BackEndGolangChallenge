package controller

import (
	"net/http"
)

type BitcoinController struct {
}

// BitcoinDetailsHandler manipula solicitações para /details/{address}
func (bc *BitcoinController) BitcoinDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Sua lógica aqui
}

// BitcoinBalanceHandler manipula solicitações para /balance/{address}
func (bc *BitcoinController) BitcoinBalanceHandler(w http.ResponseWriter, r *http.Request) {
}

// BitcoinSendHandler manipula solicitações para /send
func (bc *BitcoinController) BitcoinSendHandler(w http.ResponseWriter, r *http.Request) {
}

// GetTransactionDetails manipula solicitações para /tx/{tx}
func (bc *BitcoinController) GetTransactionDetails(w http.ResponseWriter, r *http.Request) {
}
