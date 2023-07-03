package inmemoryStorage

import (
	"repo/internal/domain/entity"
	"testing"
)

var storage *Database

func init() {
	storage = MakeDatabase()
}

func TestDatabase_AddUser(t *testing.T) {

	fedor := entity.MakeUser(37, "Fedor", "Dorofeyev", "fedor.dorofeyev@freedompay.money", "123")

	fedor, err := storage.AddUser(fedor)

	if err != nil {
		t.Errorf("error occured while add user")
	}

	if fedor.Id != 0 {
		t.Errorf("user id incorrect")
	}

	andry := entity.MakeUser(26, "Andry", "Simens", "andry.simens@freedompay.money", "123")

	andry, err = storage.AddUser(andry)

	if err != nil {
		t.Errorf("error occured while add user")
	}

	if andry.Id != 1 {
		t.Errorf("user id incorrect while add multiplie users")
	}
}

func TestDatabase_GetOne(t *testing.T) {
	user, err := storage.GetOneUser(0)

	if err != nil {
		t.Errorf("error occured while GetOneUser")
	}

	if user.Id != 0 {
		t.Error("Incorrect user id in GetOneUser method")
	}
	if user.Name != "Fedor" {
		t.Error("Incorrect user name in GetOneUser method")
	}
}

func TestDatabase_RemoveUser(t *testing.T) {
	user, err := storage.GetOneUser(0)

	if err != nil {
		t.Errorf("error occured while GetOneUser")
	}

	err = storage.RemoveUser(user)

	if err != nil {
		t.Errorf("error occured while GetOneUser")
	}

	user2, err := storage.GetOneUser(1)

	if err != nil {
		t.Errorf("error occured while GetOneUser")
	}

	err = storage.RemoveUser(user2)

	if err != nil {
		t.Errorf("error occured while GetOneUser")
	}
}

func TestDatabase_CompanyAccount(t *testing.T) {
	var err error
	accountFreedomHolding := entity.MakeCompanyAccount("аккаунт холдинга freedom", "freedom_holding", "Холдинг фридом")
	accountChocoHolding := entity.MakeCompanyAccount("аккаунт холдинга choco", "choco_holding", "Холдинг чокофемили")

	accountFreedomHolding, err = storage.AddCompanyAccount(accountFreedomHolding)
	if err != nil {
		t.Error("AddCompanyAccount error")
	}

	accountChocoHolding, err = storage.AddCompanyAccount(accountChocoHolding)
	if err != nil {
		t.Error("AddCompanyAccount error")
	}

	accountFreedomHolding.Code = "freedom_holding2"

	accountFreedomHolding, err = storage.UpdateCompanyAccount(accountFreedomHolding)
	if err != nil {
		t.Error("UpdateCompanyAccount error")
	}

	err = storage.RemoveCompanyAccount(accountFreedomHolding)
	if err != nil {
		t.Error("RemoveCompanyAccount error")
	}

	err = storage.RemoveCompanyAccount(accountChocoHolding)
	if err != nil {
		t.Error("RemoveCompanyAccount error")
	}
}

func TestDatabase_Transaction(t *testing.T) {
	accountFreedomHolding := entity.MakeCompanyAccount("аккаунт холдинга freedom", "freedom_holding", "Холдинг фридом")

	accountFreedomHolding, err := storage.AddCompanyAccount(accountFreedomHolding)

	freedomPayCompany := entity.MakeCompany(accountFreedomHolding, "freedom pay", "freedom_pay", true)
	aviataCompany := entity.MakeCompany(accountFreedomHolding, "Aviata", "aviata", true)

	freedomPayCompany, err = storage.AddCompany(freedomPayCompany)
	if err != nil {
		t.Error("error add")
	}
	freedomPayCompany, err = storage.AddCompany(aviataCompany)
	if err != nil {
		t.Error("error add")
	}

	fedor := entity.MakeUser(37, "Fedor", "Dorofeyev", "fedor.dorofeyev@freedompay.money", "123pass")

	fedor, err = storage.AddUser(fedor)
	if err != nil {
		t.Error("error add user")
	}

	FPTransactionCustomer := entity.MakeCustomer(fedor, freedomPayCompany, true, false)

	FPTransactionCustomer, err = storage.AddCustomer(FPTransactionCustomer)

	if err != nil {
		t.Error("error add customer")
	}

	FPTransactionCustomer, err = storage.UpdateCustomer(FPTransactionCustomer)
	if err != nil {
		t.Error("error update customer")
	}

	freedomPayCompanyBankAccount := entity.MakeCompanyBankAccount(freedomPayCompany,"KZ586215465431",1000000.0)
	freedomPayCompanyBankAccount, err = storage.AddCompanyBankAccount(freedomPayCompanyBankAccount)
	if err != nil {
		t.Error("error create bank account")
	}
	aviataCompanyBankAccount := entity.MakeCompanyBankAccount(aviataCompany,"KZ654613354353",0.0)
	aviataCompanyBankAccount, err = storage.AddCompanyBankAccount(aviataCompanyBankAccount)
	if err != nil {
		t.Error("error create bank account")
	}

	if !FPTransactionCustomer.CanPerformTransactions {
		t.Error("customer not permitted to perform transactions")
	} else {
		freedomPayToAviataTransfer := entity.MakeTransaction(FPTransactionCustomer,freedomPayCompanyBankAccount, aviataCompanyBankAccount,100000.0, entity.TRANSACTION_TYPE_PAY)

		freedomPayToAviataTransfer , err = storage.AddTransaction(freedomPayToAviataTransfer)
		if err != nil {
			t.Error("error create transaction")
		}

		freedomPayToAviataTransfer,err = storage.CommitTransaction(freedomPayToAviataTransfer)
		if err != nil {
			t.Error("error commit transaction")
		}

	}

}
