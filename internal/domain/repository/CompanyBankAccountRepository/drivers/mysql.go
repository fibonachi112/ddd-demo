package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MysqlCompanyBankAccountRepository struct {
	Config *config.Configuration
	Di *DI.ServiceContainer
}

func (m MysqlCompanyBankAccountRepository) GetOne(id int) (entity.CompanyBankAccount, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyBankAccountRepository) First() (entity.CompanyBankAccount, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyBankAccountRepository) SaveOne(user *entity.CompanyBankAccount) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyBankAccountRepository) UpdateOne(user *entity.CompanyBankAccount)  error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyBankAccountRepository) DeleteOne(user *entity.CompanyBankAccount) error {
	//TODO implement me
	panic("implement me")
}

func MakeMysqlDriver(config *config.Configuration, di *DI.ServiceContainer) MysqlCompanyBankAccountRepository {
	return MysqlCompanyBankAccountRepository{
		Config: config,
		Di : di,
	}
}

