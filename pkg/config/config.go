package config

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// RestyClientConfig é a configuração para o cliente Resty
var RestyClientConfig = resty.New().SetTimeout(10 * time.Second)

// RestyClient é o cliente Resty configurado
var RestyClient = RestyClientConfig

// InitRestClient inicializa e retorna um novo cliente Resty
func InitRestClient() *resty.Client {
	return resty.New()
}

// HTTPClient é um cliente HTTP configurado
var HTTPClient = &http.Client{
	Timeout: 10 * time.Second,
}

// NewRequest retorna uma nova solicitação HTTP configurada
func NewRequest() *http.Request {
	return &http.Request{
		Header: make(http.Header),
	}
}

// Init é a função que inicializa as configurações
func Init() {
	RestyClient = InitRestClient()
	HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}
