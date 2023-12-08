package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL da API Explorer
	apiURL := "https://bitcoin.explorer.klever.io/api/v2/address/bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r"

	// Credenciais para autenticação
	username := "support"
	password := "Fg+GJKDACKIEOD3XVps="

	// Codificar as credenciais para autenticação básica
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	// Configurar uma solicitação HTTP
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Erro ao criar a solicitação:", err)
		return
	}

	// Adicionar cabeçalho de autorização para autenticação básica
	req.Header.Add("Authorization", "Basic "+auth)

	// Realizar a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao realizar a solicitação:", err)
		return
	}
	defer resp.Body.Close()

	// Ler e exibir o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}
	fmt.Println(string(body))
}
