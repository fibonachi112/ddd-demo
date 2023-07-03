package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MockCompanyAccountRepository struct {
	Config *config.Configuration
	Di     *DI.ServiceContainer
}

func (r MockCompanyAccountRepository) GetOne(id int) (entity.CompanyBankAccount, error) {
	//TODO implement me
	companyBankAccount, err := r.Di.InMemoryStorage.GetOneCompanyBankAccount(id)
	if err != nil {
		return companyBankAccount, err
	}
	return companyBankAccount, nil
}


func (r MockCompanyAccountRepository) SaveOne(companyBankAccount *entity.CompanyBankAccount) error {
	//TODO implement me
	newCompanyAccount, err := r.Di.InMemoryStorage.AddCompanyBankAccount(*companyBankAccount)
	if err != nil {
		return err
	}
	companyBankAccount.Id = newCompanyAccount.Id
	companyBankAccount.CreatedAt = newCompanyAccount.CreatedAt

	return nil
}

func (r MockCompanyAccountRepository) UpdateOne(companyBankAccount *entity.CompanyBankAccount) error {
	updatedCompanyBankAccount, err := r.Di.InMemoryStorage.UpdateCompanyBankAccount(*companyBankAccount)
	if err != nil {
		return err
	}
	companyBankAccount = &updatedCompanyBankAccount
	return nil
}

func (r MockCompanyAccountRepository) DeleteOne(companyBankAccount *entity.CompanyBankAccount) error {
	err := r.Di.InMemoryStorage.RemoveCompanyBankAccount(*companyBankAccount)
	if err != nil {
		return err
	}
	return nil
}

func MakeMockDriver(config *config.Configuration, di *DI.ServiceContainer) MockCompanyAccountRepository {
	return MockCompanyAccountRepository{
		Config: config,
		Di:     di,
	}
}
