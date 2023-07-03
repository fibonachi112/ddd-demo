package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MockTransactionRepository struct {
	Config *config.Configuration
	Di     *DI.ServiceContainer
}

func (r MockTransactionRepository) GetOne(id int) (entity.Transaction, error) {
	//TODO implement me
	transaction, err := r.Di.InMemoryStorage.GetOneTransaction(id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}


func (r MockTransactionRepository) SaveOne(transaction *entity.Transaction) error {
	//TODO implement me
	newTransaction, err := r.Di.InMemoryStorage.AddTransaction(*transaction)
	if err != nil {
		return err
	}
	transaction.Id = newTransaction.Id
	transaction.CreatedAt = newTransaction.CreatedAt

	return nil
}

func (r MockTransactionRepository) UpdateOne(transaction *entity.Transaction) error {
	updatedTransaction, err := r.Di.InMemoryStorage.UpdateTransaction(*transaction)
	if err != nil {
		return err
	}
	transaction = &updatedTransaction
	return nil
}

func (r MockTransactionRepository) DeleteOne(transaction *entity.Transaction) error {
	err := r.Di.InMemoryStorage.RemoveTransaction(*transaction)
	if err != nil {
		return err
	}
	return nil
}

func (r MockTransactionRepository) Commit(transaction *entity.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func MakeMockDriver(config *config.Configuration, di *DI.ServiceContainer) MockTransactionRepository {
	return MockTransactionRepository{
		Config: config,
		Di:     di,
	}
}
