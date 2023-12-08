// backendgolangchallenge/pkg/service/bitcoin_service.go
package service

import "backendgolangchallenge/pkg/model"

type BitcoinService interface {
	SelectUTXOsForSend(address string, amount int64) (*model.BitcoinSendRequest, error)
	GetTransactionDetails(txID string) (*model.BitcoinTransactionDetails, error)
}
