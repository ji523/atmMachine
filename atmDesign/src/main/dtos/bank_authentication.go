package dtos

type BankAuthenticationRequest struct {
	PinCode int64
	CardNumber string
}

type BankAuthenticationResponse struct {
	IsValid bool
}
