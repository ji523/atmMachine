package clients

import (
	"atmDesign/src/main/config"
	"atmDesign/src/main/dtos"
)
var BankConfigVar BankConfig
type IBankClient interface {
	AuthenticateFromBank(*dtos.BankAuthenticationRequest) *dtos.BankAuthenticationResponse
}
type BankConfig struct {
	BaseUrl       string
	AuthenticationFromBankApi config.Api
}

type BankClient struct {
	BaseUrl                   string
	AuthenticationFromBankApi config.Api
}

func NewBankConfig() *BankConfig {
	return &BankConfig{
		BaseUrl: "BaseUrl",
		AuthenticationFromBankApi: config.Api{
			ApiPath:     "bankClientGetPath",
			HttpMethod:  "bankClientGetHttpMethod",
			ReadTimeout: 0,
		},
	}
}
func GetBankConfig() *BankConfig {
	return &BankConfigVar
}
