// bitcoin_transaction_details.go
package model

// BitcoinTransactionDetails representa os detalhes de uma transação de Bitcoin
type BitcoinTransactionDetails struct {
	Addresses []BitcoinTransactionAddress `json:"addresses"`
	Block     int64                       `json:"block"`
	TxID      string                      `json:"txID"`
}

// BitcoinTransactionAddress representa um endereço associado a uma transação de Bitcoin
type BitcoinTransactionAddress struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}
