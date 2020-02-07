package impl
import (
	"atmDesign/src/main/dtos"
	"atmDesign/src/main/services"
)

type TransactionServiceImpl struct {
	transactionRepo *TransactionRepo
}

func (t TransactionServiceImpl) Authenticate(request *interface{}) *interface{} {
	panic("implement me")
}

func NewTransactionServiceImpl() services.ITransactionService {
	return &TransactionServiceImpl{}
}

func (a AuthenticationServiceImpl) Withdraw(request *dtos.AuthenticationRequest) *dtos.AuthenticationResponse {
	res := a.transactionRepo.DoTransaction()
	return map_repo_to_dto_transaction(res);
}
