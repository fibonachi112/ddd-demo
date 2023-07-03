package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MysqlCustomerRepository struct {
	Config *config.Configuration
	Di *DI.ServiceContainer
}

func (m MysqlCustomerRepository) GetOne(id int) (entity.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCustomerRepository) SaveOne(customer *entity.Customer) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCustomerRepository) UpdateOne(customer *entity.Customer)  error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCustomerRepository) DeleteOne(customer *entity.Customer) error {
	//TODO implement me
	panic("implement me")
}

func MakeMysqlDriver(config *config.Configuration, di *DI.ServiceContainer) MysqlCustomerRepository {
	return MysqlCustomerRepository{
		Config: config,
		Di : di,
	}
}

