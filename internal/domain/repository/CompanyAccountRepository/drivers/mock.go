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

func (r MockCompanyAccountRepository) GetOne(id int) (entity.CompanyAccount, error) {
	//TODO implement me
	companyAccount, err := r.Di.InMemoryStorage.GetOneCompanyAccount(id)
	if err != nil {
		return companyAccount, err
	}
	return companyAccount, nil
}


func (r MockCompanyAccountRepository) SaveOne(companyAccount *entity.CompanyAccount) error {
	//TODO implement me
	newCompanyAccount, err := r.Di.InMemoryStorage.AddCompanyAccount(*companyAccount)
	if err != nil {
		return err
	}
	companyAccount.Id = newCompanyAccount.Id
	companyAccount.CreatedAt = newCompanyAccount.CreatedAt

	return nil
}

func (r MockCompanyAccountRepository) UpdateOne(companyAccount *entity.CompanyAccount) error {
	updatedCompanyAccount, err := r.Di.InMemoryStorage.UpdateCompanyAccount(*companyAccount)
	if err != nil {
		return err
	}
	companyAccount = &updatedCompanyAccount
	return nil
}

func (r MockCompanyAccountRepository) DeleteOne(companyAccount *entity.CompanyAccount) error {
	err := r.Di.InMemoryStorage.RemoveCompanyAccount(*companyAccount)
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
