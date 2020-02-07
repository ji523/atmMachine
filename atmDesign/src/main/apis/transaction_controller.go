package apis

import (
	"atmDesign/src/main/services"
	"net/http"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *TransactionController {
	return &TransactionController{transactionService:transactionService};
}

func(controller *AuthenticationController) authenticate(req *http.Request) *Response {
	return controller.transactionService.Authenticate()
}