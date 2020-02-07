package services

import "atmDesign/src/main/dtos"

type IAuthenticationService interface {
	Authenticate(request *dtos.AuthenticationRequest) *dtos.AuthenticationResponse
}
