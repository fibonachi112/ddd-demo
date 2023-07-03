package CompanyAccountRepository

import (
	"repo/internal/domain/entity"
)

type ICompanyAccountRepository interface {
	GetOne(id int) (entity.CompanyAccount, error)
	SaveOne(account *entity.CompanyAccount) error
	UpdateOne(account *entity.CompanyAccount) error
	DeleteOne(account *entity.CompanyAccount) error
}
