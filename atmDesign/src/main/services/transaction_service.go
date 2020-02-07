package services

import "atmDesign/src/main/dtos"

type ITransactionService interface {
	Authenticate(request *dtos.TransactionRequest) *dtos.TransactionResponse
}
