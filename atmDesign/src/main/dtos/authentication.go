package dtos

type AuthenticationRequest struct {
	PinCode int64
	CardNumber string
}

type AuthenticationResponse struct {
	IsValid bool
}
