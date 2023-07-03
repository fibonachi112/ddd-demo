package entity

import "time"

const TRANSACTION_TYPE_FILL = "fill"

const TRANSACTION_TYPE_PAY = "pay"

type Transaction struct {
	Id                     int
	Type                   string
	Sum                    float64
	InitiatorCustomer      Customer
	CompanyBankAccountFrom CompanyBankAccount
	CompanyBankAccountTo   CompanyBankAccount
	CreatedAt              string
	UpdatedAt              string
}

func (t *Transaction) SetCreatedAtNow() {
	t.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}

func (t *Transaction) SetUpdatedAtNow() {
	t.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
}

func MakeTransaction(customer Customer, accountFrom, accountTo CompanyBankAccount, sum float64, transactionType string) Transaction {
	t := Transaction{
		Id:                     0,
		Type:                   transactionType,
		Sum:                    sum,
		InitiatorCustomer:      customer,
		CompanyBankAccountFrom: accountFrom,
		CompanyBankAccountTo:   accountTo,
		CreatedAt:              "",
		UpdatedAt:              "",
	}
	t.SetCreatedAtNow()
	return t
}
