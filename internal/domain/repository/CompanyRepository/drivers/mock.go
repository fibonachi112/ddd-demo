package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MockCompanyRepository struct {
	Config *config.Configuration
	Di     *DI.ServiceContainer
}

func (r MockCompanyRepository) GetOne(id int) (entity.Company, error) {
	//TODO implement me
	company, err := r.Di.InMemoryStorage.GetOneCompany(id)
	if err != nil {
		return company, err
	}
	return company, nil
}


func (r MockCompanyRepository) SaveOne(company *entity.Company) error {
	//TODO implement me
	newCompany, err := r.Di.InMemoryStorage.AddCompany(*company)
	if err != nil {
		return err
	}
	company.Id = newCompany.Id
	company.CreatedAt = newCompany.CreatedAt

	return nil
}

func (r MockCompanyRepository) UpdateOne(company *entity.Company) error {
	updatedCompany, err := r.Di.InMemoryStorage.UpdateCompany(*company)
	if err != nil {
		return err
	}
	company = &updatedCompany
	return nil
}

func (r MockCompanyRepository) DeleteOne(company *entity.Company) error {
	err := r.Di.InMemoryStorage.RemoveCompany(*company)
	if err != nil {
		return err
	}
	return nil
}

func MakeMockDriver(config *config.Configuration, di *DI.ServiceContainer) MockCompanyRepository {
	return MockCompanyRepository{
		Config: config,
		Di:     di,
	}
}
