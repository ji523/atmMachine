package config

import (
	"atmDesign/src/main/clients"
	"net/http"
"time"
)

type Api struct {
	ApiPath     string
	HttpMethod  string
	ReadTimeout int32
	QueryParams []string
	PathParams  []string
}

func createHTTPClient(requestTimeout int32) *http.Client {
	client := &http.Client{
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
	return client
}

func SetUpClients() {
	setUpBankClient()
}

func setUpBankClient() {
	clients.BankConfigVar = *clients.NewBankConfig()
}

func BankHttpClient() clients.HttpClient {
	return createHTTPClient(clients.BankConfigVar.AuthenticationFromBankApi.ReadTimeout)
}

