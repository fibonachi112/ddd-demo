package CustomerRepository

import (
	"repo/internal/domain/entity"
)

type ICustomerRepository interface {
	GetOne(id int) (entity.Customer, error)
	SaveOne(company *entity.Customer) error
	UpdateOne(company *entity.Customer) error
	DeleteOne(company *entity.Customer) error
}
