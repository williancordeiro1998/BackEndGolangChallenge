// bitcoin_address_details.go
package model

// AddressDetails representa os detalhes de um endereço Bitcoin
type AddressDetails struct {
	Address string       `json:"address"`
	Balance Balance      `json:"balance"`
	TotalTx int          `json:"totalTx"`
	Total   TotalDetails `json:"total"`
}

// Balance representa o saldo de um endereço Bitcoin
type Balance struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

// TotalDetails contém detalhes sobre transações totais
type TotalDetails struct {
	Sent     string `json:"sent"`
	Received string `json:"received"`
}
