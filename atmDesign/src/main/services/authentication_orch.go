package services

import "atmDesign/src/main/dtos"

type IAuthenticationOrch interface {
	Authenticate(*dtos.AuthenticationRequest) *dtos.AuthenticationResponse
}
