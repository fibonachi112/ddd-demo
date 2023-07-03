package CompanyBankAccountRepository

import (
	"repo/internal/domain/entity"
)

type ICompanyBankAccountRepository interface {
	GetOne(id int) (entity.CompanyBankAccount, error)
	SaveOne(account *entity.CompanyBankAccount) error
	UpdateOne(account *entity.CompanyBankAccount) error
	DeleteOne(account *entity.CompanyBankAccount) error
}
