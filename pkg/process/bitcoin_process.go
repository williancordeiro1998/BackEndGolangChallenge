package process

import (
	"io/ioutil"
	"net/http"
)

// ProcessData é a função que processa os dados
func ProcessData(w http.ResponseWriter, r *http.Request) {
	// URL da API Blockbook
	url := "https://bitcoin.blockbook.chains.klever.io/api/v2"

	// Faz uma solicitação GET para a URL da API Blockbook
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erro ao fazer solicitação para a API Blockbook: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Verifica se o status da resposta é 200 OK
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "A API Blockbook retornou um status de erro: "+resp.Status, http.StatusInternalServerError)
		return
	}

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a resposta da API Blockbook: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Escreve o corpo da resposta para o ResponseWriter
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Erro ao escrever a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
