// bitcoin_send_request.go
package model

// BitcoinSendRequest representa a solicitação de envio de Bitcoin
type BitcoinSendRequest struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

// BitcoinSendResponse representa a resposta do envio de Bitcoin
type BitcoinSendResponse struct {
	UTXOs []BitcoinUTXO `json:"utxos"`
}

// BitcoinUTXO representa uma saída de transação não gasta (UTXO)
type BitcoinUTXO struct {
	TxID   string `json:"txid"`
	Amount string `json:"amount"`
}
