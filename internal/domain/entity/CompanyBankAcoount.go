package entity

import "time"

// Сущность банковского счёта компании
type CompanyBankAccount struct {
	Id        int
	Company   Company
	Number    string
	Sum       float64
	CreatedAt string
	UpdatedAt string
}

func (c *CompanyBankAccount) SetCreatedAtNow() {
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}
func (c *CompanyBankAccount) SetUpdatedAtNow() {
	c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
}

func MakeCompanyBankAccount(company Company, number string, sum float64) CompanyBankAccount {
	finAccount := CompanyBankAccount{
		Id:        0,
		Company:   company,
		Number:    number,
		Sum:       sum,
		CreatedAt: "",
		UpdatedAt: "",
	}

	finAccount.SetCreatedAtNow()
	return finAccount
}
