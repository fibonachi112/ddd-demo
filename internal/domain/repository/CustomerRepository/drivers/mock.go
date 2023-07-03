package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MockCustomerRepository struct {
	Config *config.Configuration
	Di     *DI.ServiceContainer
}

func (r MockCustomerRepository) GetOne(id int) (entity.Customer, error) {
	//TODO implement me
	customer, err := r.Di.InMemoryStorage.GetOneCustomer(id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}


func (r MockCustomerRepository) SaveOne(customer *entity.Customer) error {
	//TODO implement me
	newCustomer, err := r.Di.InMemoryStorage.AddCustomer(*customer)
	if err != nil {
		return err
	}
	customer.Id = newCustomer.Id
	customer.CreatedAt = newCustomer.CreatedAt

	return nil
}

func (r MockCustomerRepository) UpdateOne(customer *entity.Customer) error {
	updatedCustomer, err := r.Di.InMemoryStorage.UpdateCustomer(*customer)
	if err != nil {
		return err
	}
	customer = &updatedCustomer
	return nil
}

func (r MockCustomerRepository) DeleteOne(customer *entity.Customer) error {
	err := r.Di.InMemoryStorage.RemoveCustomer(*customer)
	if err != nil {
		return err
	}
	return nil
}

func MakeMockDriver(config *config.Configuration, di *DI.ServiceContainer) MockCustomerRepository {
	return MockCustomerRepository{
		Config: config,
		Di:     di,
	}
}
