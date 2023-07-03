package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MysqlCompanyAccountRepository struct {
	Config *config.Configuration
	Di *DI.ServiceContainer
}

func (m MysqlCompanyAccountRepository) GetOne(id int) (entity.CompanyAccount, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyAccountRepository) First() (entity.CompanyAccount, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyAccountRepository) SaveOne(user *entity.CompanyAccount) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyAccountRepository) UpdateOne(user *entity.CompanyAccount)  error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyAccountRepository) DeleteOne(user *entity.CompanyAccount) error {
	//TODO implement me
	panic("implement me")
}

func MakeMysqlDriver(config *config.Configuration, di *DI.ServiceContainer) MysqlCompanyAccountRepository {
	return MysqlCompanyAccountRepository{
		Config: config,
		Di : di,
	}
}

