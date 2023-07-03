package entity

import "time"

// Сущность аккаунта компании
type CompanyAccount struct {
	Id          int
	Name        string
	Code        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

func (c *CompanyAccount) SetCreatedAtNow(){
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}
func (c *CompanyAccount) SetUpdatedAtNow(){
	c.UpdatedAt= time.Now().Format("2006-01-02 15:04:05")
}



func MakeCompanyAccount(name, code, description string) CompanyAccount {
	return CompanyAccount{
		Id:          0,
		Name:        name,
		Code:        code,
		Description: description,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   "",
	}
}
