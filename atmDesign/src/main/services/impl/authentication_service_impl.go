package impl

import (
	"atmDesign/src/main/dtos"
	"atmDesign/src/main/services"
)

type AuthenticationServiceImpl struct {
}

func NewAuthenticationServiceImpl() services.IAuthenticationService {
	return &AuthenticationServiceImpl{}
}

func (a AuthenticationServiceImpl) Authenticate(request *dtos.AuthenticationRequest) *dtos.AuthenticationResponse {
	cardNumber := request.CardNumber
	pinCode := request.PinCode
	if pinCode < 1000 || len(cardNumber) < 8 {
		return &dtos.AuthenticationResponse{IsValid:false}
	}
	return &dtos.AuthenticationResponse{
		IsValid:true,
	}
}


