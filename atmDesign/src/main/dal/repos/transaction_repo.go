package repos

type TransactionRepo struct {
	dbConn *gorm.Db
}

func NewTransactionRepo(dbConn *gorm.Db) *TransactionRepo{
	return &TransactionRepo{dbConn:nil}
}