package entity

import "time"

// Сущность компании
type Company struct {
	Id         int
	Account    CompanyAccount
	Name       string
	Code       string
	Active    bool
	CreatedAt string
	UpdatedAt string
}

func (c *Company) SetCreatedAtNow(){
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}
func (c *Company) SetUpdatedAtNow(){
	c.UpdatedAt= time.Now().Format("2006-01-02 15:04:05")
}



func MakeCompany(account CompanyAccount, name, code string, active bool) Company {
	return Company{
		Id:        0,
		Account:   account,
		Name:      name,
		Code:      code,
		Active:    active,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: "",
	}
}
