// bitcoin_utils.go
package util

import (
	"math/big"
)

// ConvertSatoshisToBTC converte satoshis para BTC
func ConvertSatoshisToBTC(satoshis *big.Int) *big.Int {
	return new(big.Int).Div(satoshis, big.NewInt(100000000))
}

// ConvertBTCtoSatoshis converte BTC para satoshis
func ConvertBTCtoSatoshis(btc *big.Int) *big.Int {
	return new(big.Int).Mul(btc, big.NewInt(100000000))
}
