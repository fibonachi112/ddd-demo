package repository

import (
	"repo/config"
	DiFabric "repo/internal/DI/fabric"
	"repo/internal/domain/entity"
	"testing"
)

func TestTransactions(t *testing.T) {
	var err error

	di := DiFabric.Make()
	config.Config.Env = config.ENV_TEST




	// создаем аккаунт компании
	accountFreedomHolding := entity.MakeCompanyAccount(
		"Аккаунт холдинга freedom",
		"freedom_holding",
		"Холдинг фридом")

	err = di.CompanyAccountRepository.SaveOne(&accountFreedomHolding)
	if err != nil {
		t.Error("error while creating company account")
	}


	// создаем компанию freedom pay и соотносим ее с холдингом фридома
	freedomPayCompany := entity.MakeCompany(
		accountFreedomHolding,
		"freedom pay","freedom_pay",true)

	err = di.CompanyRepository.SaveOne(&freedomPayCompany)
	if err != nil {
		t.Error("error while creating company")
	}


	// создаем компанию aviata и соотносим ее с холдингом фридома
	aviataCompany := entity.MakeCompany(
		accountFreedomHolding,"Aviata","aviata",true)

	err = di.CompanyRepository.SaveOne(&aviataCompany)
	if err != nil {
		t.Error("error while creating company")
	}


	// создаем пользователя от лица которого будут производиться операции
	fedor := entity.MakeUser(
		37,"Fedor","Dorofeyev",
		"fedor.dorofeyev@gmail.com","mgfh{tr52E%}2Wdf")

	err = di.UserRepository.SaveOne(&fedor)
	if err != nil {
		t.Error("Failed on user create")
	}


	// далем пользователя клиентом зарегистрированной у нас компании
	// даем ему право проводить финансовые транзакции
	FPTransactionCustomer := entity.MakeCustomer(
		fedor,freedomPayCompany,
		true,false)


	err = di.CustomerRepository.SaveOne(&FPTransactionCustomer)
	if err != nil {
		t.Error("Failed on create customer")
	}

	// заводим счёт с положительным балансом компании freedom pay
	freedomPayCompanyBankAccount := entity.MakeCompanyBankAccount(
		freedomPayCompany, "KZ586215465431", 1000000000.0)

	err = di.CompanyBankAccountRepository.SaveOne(&freedomPayCompanyBankAccount)
	if err != nil {
		t.Error("Failed on create company bank account")
	}

	// заводим счёт компании авиата
	aviataCompanyBankAccount := entity.MakeCompanyBankAccount(
		aviataCompany, "KZ654613354353", 0.0)

	err = di.CompanyBankAccountRepository.SaveOne(&aviataCompanyBankAccount)
	if err != nil {
		t.Error("Failed on create company bank account")
	}





	//производим транзакцию с переводом одной компании в другую от лица кастомера
	if !FPTransactionCustomer.CanPerformTransactions {
		t.Error("customer not permitted to perform transactions")
	} else {
		// создаем транзакцию с платежём со счёта freedom pay на с чёт aviata
		freedomPayToAviataTransfer := entity.MakeTransaction(
			FPTransactionCustomer, freedomPayCompanyBankAccount, aviataCompanyBankAccount,
			10000000.0, entity.TRANSACTION_TYPE_PAY)

		err = di.TransactionRepository.SaveOne(&freedomPayToAviataTransfer)
		if err != nil {
			t.Error("Failed on create transaction")
		}

		err = di.TransactionRepository.Commit(&freedomPayToAviataTransfer)
		if err != nil {
			t.Error("Failed on commit transaction")
		}
	}
}


