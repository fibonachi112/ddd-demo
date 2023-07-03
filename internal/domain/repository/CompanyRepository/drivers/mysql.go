package drivers

import (
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MysqlCompanyRepository struct {
	Config *config.Configuration
	Di *DI.ServiceContainer
}

func (m MysqlCompanyRepository) GetOne(id int) (entity.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyRepository) SaveOne(company *entity.Company) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyRepository) UpdateOne(company *entity.Company)  error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlCompanyRepository) DeleteOne(company *entity.Company) error {
	//TODO implement me
	panic("implement me")
}

func MakeMysqlDriver(config *config.Configuration, di *DI.ServiceContainer) MysqlCompanyRepository {
	return MysqlCompanyRepository{
		Config: config,
		Di : di,
	}
}

