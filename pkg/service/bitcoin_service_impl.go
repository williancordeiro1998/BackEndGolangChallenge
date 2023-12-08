package service

import "backendgolangchallenge/pkg/model"

type BitcoinServiceImpl struct {
}

// SelectUTXOsForSend implementa o método da interface BitcoinService
func (s *BitcoinServiceImpl) SelectUTXOsForSend(address string, amount int64) (*model.BitcoinSendRequest, error) {
	return &model.BitcoinSendRequest{
		Address: address,
		Amount:  "10",
	}, nil
}

// GetTransactionDetails implementa o método da interface BitcoinService
func (s *BitcoinServiceImpl) GetTransactionDetails(txID string) (*model.BitcoinTransactionDetails, error) {
	return &model.BitcoinTransactionDetails{}, nil
}
