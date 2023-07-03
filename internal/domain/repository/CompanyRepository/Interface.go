package CompanyRepository

import (
	"repo/internal/domain/entity"
)

type ICompanyRepository interface {
	GetOne(id int) (entity.Company, error)
	SaveOne(company *entity.Company) error
	UpdateOne(company *entity.Company) error
	DeleteOne(company *entity.Company) error
}
