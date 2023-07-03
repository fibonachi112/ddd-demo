package TransactionRepository

import (
	"repo/internal/domain/entity"
)

type ITransactionRepository interface {
	GetOne(id int) (entity.Transaction, error)
	SaveOne(transaction *entity.Transaction) error
	UpdateOne(transaction *entity.Transaction) error
	DeleteOne(transaction *entity.Transaction) error
	Commit(transaction *entity.Transaction) error
}
