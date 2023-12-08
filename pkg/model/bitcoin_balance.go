// bitcoin_balance.go
package model

// BitcoinBalance representa o saldo Bitcoin
type BitcoinBalance struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}
