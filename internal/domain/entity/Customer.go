package entity

import (
	"time"
)

// Сущность кастомера
type Customer struct {
	Id                     int
	User                   User
	Company                Company
	CanPerformTransactions bool
	Transactions           []Transaction
	IsAccountAdmin         bool
	CreatedAt              string
	UpdatedAt              string
}

func (c *Customer) SetCreatedAtNow() {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}

func (c *Customer) SetUpdatedAtNow() {
	c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
}

func MakeCustomer(user User, company Company, canPerformTransactions, isAccountAdmin bool) Customer {
	return Customer{
		Id:                     0,
		User:                   user,
		Company:                company,
		CanPerformTransactions: canPerformTransactions,
		IsAccountAdmin:         isAccountAdmin,
		CreatedAt:              time.Now().Format("2006-01-02 15:04:05"),
	}
}
