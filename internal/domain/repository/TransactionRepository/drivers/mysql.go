package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MysqlTransactionRepository struct {
	Config *config.Configuration
	Di *DI.ServiceContainer
}

func (m MysqlTransactionRepository) GetOne(id int) (entity.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlTransactionRepository) SaveOne(Transaction *entity.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlTransactionRepository) UpdateOne(Transaction *entity.Transaction)  error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlTransactionRepository) DeleteOne(Transaction *entity.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlTransactionRepository) Commit(transaction *entity.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func MakeMysqlDriver(config *config.Configuration, di *DI.ServiceContainer) MysqlTransactionRepository {
	return MysqlTransactionRepository{
		Config: config,
		Di : di,
	}
}

