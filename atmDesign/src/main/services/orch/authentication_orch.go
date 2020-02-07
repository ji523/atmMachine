package orch

import (
	"atmDesign/src/main/clients"
	"atmDesign/src/main/dtos"
	"atmDesign/src/main/services"
)

type AuthenticationOrch struct {
	authenticationService services.IAuthenticationService
	bankClient clients.IBankClient
}

func NewAuthenticationOrch(authenticationService services.IAuthenticationService,
	bankClient clients.IBankClient) services.IAuthenticationOrch {
		if authenticationService == nil || bankClient == nil {
			panic("dependency is not provided")
		}
	return &AuthenticationOrch{authenticationService:authenticationService, bankClient:bankClient}
}


func (orch *AuthenticationOrch) Authenticate(req *dtos.AuthenticationRequest) *dtos.AuthenticationResponse {
	isValid := orch.authenticationService.Authenticate(req).IsValid
	if isValid {
		bankReq := &dtos.BankAuthenticationRequest{PinCode:req.PinCode, CardNumber:req.CardNumber}
		isValid = orch.bankClient.AuthenticateFromBank(bankReq).IsValid
	}
	return &dtos.AuthenticationResponse{IsValid:isValid}
}
