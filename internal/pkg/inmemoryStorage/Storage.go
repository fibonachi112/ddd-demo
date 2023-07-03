package inmemoryStorage

import (
	"errors"
	"repo/internal/domain/entity"
	"sync"
)

var DB *Database

var dbOnce sync.Once

type Database struct {
	mx                  sync.Mutex
	Users               []entity.User
	CompanyAccounts     []entity.CompanyAccount
	Companies           []entity.Company
	Customers           []entity.Customer
	CompanyBankAccounts []entity.CompanyBankAccount
	Transactions        []entity.Transaction
}

func init(){
	dbOnce.Do(func() {
		DB = MakeDatabase()
	})
}

func MakeDatabase() *Database {
	return &Database{
		mx:                  sync.Mutex{},
		Users:               make([]entity.User, 0),
		CompanyAccounts:     make([]entity.CompanyAccount, 0),
		Companies:           make([]entity.Company, 0),
		Customers:           make([]entity.Customer, 0),
		CompanyBankAccounts: make([]entity.CompanyBankAccount, 0),
	}
}

func (db *Database) AddUser(user entity.User) (entity.User, error) {
	db.mx.Lock()
	user.Id = len(db.Users)
	db.Users = append(db.Users, user)
	db.mx.Unlock()
	return user, nil
}

func (db *Database) GetOneUser(id int) (entity.User, error) {
	for _, u := range db.Users {
		if u.Id == id {
			return u, nil
		}
	}
	return entity.User{}, errors.New("User not found")
}

func (db *Database) UpdateUser(user entity.User) (entity.User, error) {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, u := range db.Users {
		if u.Id == user.Id {
			user.SetUpdatedAtNow()
			db.Users[i] = user
			return user, nil
		}
	}

	return entity.User{}, errors.New("User not found")
}

func (db *Database) RemoveUser(user entity.User) error {
	db.mx.Lock()
	defer db.mx.Unlock()

	for i, u := range db.Users {
		if u.Id == user.Id {
			if len(db.Users) > 1 {
				db.Users = append(db.Users[0:i], db.Users[i+1:]...)
				return nil
			} else {
				db.Users = make([]entity.User, 0)
				return nil
			}
		}
	}

	return errors.New("user not found")
}

func (db *Database) AddCompanyAccount(account entity.CompanyAccount) (entity.CompanyAccount, error) {
	db.mx.Lock()
	account.Id = len(db.CompanyAccounts)
	db.CompanyAccounts = append(db.CompanyAccounts, account)
	db.mx.Unlock()

	return account, nil
}

func (db *Database) GetOneCompanyAccount(id int) (entity.CompanyAccount, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	for _, a := range db.CompanyAccounts {
		if a.Id == id {
			return a, nil
		}
	}

	return entity.CompanyAccount{}, errors.New("company not found")
}

func (db *Database) UpdateCompanyAccount(account entity.CompanyAccount) (entity.CompanyAccount, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	for i, u := range db.CompanyAccounts {
		if u.Id == account.Id {
			account.SetUpdatedAtNow()
			db.CompanyAccounts[i] = account
			return account, nil
		}
	}

	return entity.CompanyAccount{}, errors.New("account not found")
}

func (db *Database) RemoveCompanyAccount(account entity.CompanyAccount) error {
	db.mx.Lock()
	defer db.mx.Unlock()

	for i, a := range db.CompanyAccounts {
		if a.Id == account.Id {
			if len(db.CompanyAccounts) > 1 {
				db.CompanyAccounts = append(db.CompanyAccounts[0:i], db.CompanyAccounts[i+1:]...)
				return nil
			} else {
				db.CompanyAccounts = make([]entity.CompanyAccount, 0)
				return nil
			}
		}
	}

	return errors.New("company account not found")
}

func (db *Database) AddCompany(company entity.Company) (entity.Company, error) {
	company.Id = len(db.Companies)
	_, err := db.GetOneCompanyAccount(company.Account.Id)
	if err != nil {
		return entity.Company{}, errors.New("account must exist")
	}
	db.mx.Lock()
	defer db.mx.Unlock()
	db.Companies = append(db.Companies, company)
	return company, nil
}

func (db *Database) GetOneCompany(id int) (entity.Company, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	for _, c := range db.Companies {
		if c.Id == id {
			return c, nil
		}
	}
	return entity.Company{}, errors.New("Company not found")
}

func (db *Database) UpdateCompany(company entity.Company) (entity.Company, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	for i, u := range db.Companies {
		if u.Id == company.Id {
			company.SetUpdatedAtNow()
			db.Companies[i] = company
			return company, nil
		}
	}
	return entity.Company{}, errors.New("Company not found")
}

func (db *Database) RemoveCompany(company entity.Company) error {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, c := range db.Companies {
		if c.Id == company.Id {
			if len(db.Companies) > 1 {
				db.Companies = append(db.Companies[0:i], db.Companies[i+1:]...)
				return nil
			} else {
				db.Companies = make([]entity.Company, 0)
				return nil
			}
		}
	}
	return errors.New("company not found")
}

func (db *Database) AddCustomer(customer entity.Customer) (entity.Customer, error) {
	customer.Id = len(db.Customers)
	_, err := db.GetOneCompany(customer.Company.Id)
	if err != nil {
		return entity.Customer{}, errors.New("company must exist")
	}
	db.mx.Lock()
	defer db.mx.Unlock()
	db.Customers = append(db.Customers, customer)
	return customer, nil
}

func (db *Database) GetOneCustomer(id int) (entity.Customer, error) {
	db.mx.Lock()
	defer db.mx.Unlock()
	for _, c := range db.Customers {
		if c.Id == id {
			return c, nil
		}
	}
	return entity.Customer{}, errors.New("Customer not found")
}

func (db *Database) UpdateCustomer(customer entity.Customer) (entity.Customer, error) {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, u := range db.Customers {
		if u.Id == customer.Id {
			customer.SetUpdatedAtNow()
			db.Customers[i] = customer
			return customer, nil
		}
	}
	return entity.Customer{}, errors.New("customer not found")
}

func (db *Database) RemoveCustomer(customer entity.Customer) error {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, c := range db.Customers {
		if c.Id == customer.Id {
			if len(db.Customers) > 1 {
				db.Customers = append(db.Customers[0:i], db.Customers[i+1:]...)
				return nil
			} else {
				db.Customers = make([]entity.Customer, 0)
				return nil
			}
		}
	}
	return errors.New("customer not found")
}

func (db *Database) AddCompanyBankAccount(bankAccount entity.CompanyBankAccount) (entity.CompanyBankAccount, error) {
	bankAccount.Id = len(db.CompanyBankAccounts)
	_, err := db.GetOneCompany(bankAccount.Company.Id)
	if err != nil {
		return entity.CompanyBankAccount{}, errors.New("company must exist")
	}
	db.mx.Lock()
	defer db.mx.Unlock()
	db.CompanyBankAccounts = append(db.CompanyBankAccounts, bankAccount)
	return bankAccount, nil
}

func (db *Database) GetOneCompanyBankAccount(id int) (entity.CompanyBankAccount, error) {
	db.mx.Lock()
	defer db.mx.Unlock()
	for _, c := range db.CompanyBankAccounts {
		if c.Id == id {
			return c, nil
		}
	}
	return entity.CompanyBankAccount{}, errors.New("CompanyBankAccount not found")
}

func (db *Database) UpdateCompanyBankAccount(bankAccount entity.CompanyBankAccount) (entity.CompanyBankAccount, error) {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, u := range db.CompanyBankAccounts {
		if u.Id == bankAccount.Id {
			bankAccount.SetUpdatedAtNow()
			db.CompanyBankAccounts[i] = bankAccount
			return bankAccount, nil
		}
	}
	return entity.CompanyBankAccount{}, errors.New("bankAccount not found")
}

func (db *Database) RemoveCompanyBankAccount(bankAccount entity.CompanyBankAccount) error {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, c := range db.CompanyBankAccounts {
		if c.Id == bankAccount.Id {
			if len(db.CompanyBankAccounts) > 1 {
				db.CompanyBankAccounts = append(db.CompanyBankAccounts[0:i], db.CompanyBankAccounts[i+1:]...)
				return nil
			} else {
				db.CompanyBankAccounts = make([]entity.CompanyBankAccount, 0)
				return nil
			}
		}
	}
	return errors.New("bankAccount not found")
}

func (db *Database) AddTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	transaction.Id = len(db.Companies)
	_, err := db.GetOneCompanyBankAccount(transaction.CompanyBankAccountFrom.Id)
	if err != nil {
		return entity.Transaction{}, errors.New("bank account must exist")
	}
	_, err = db.GetOneCompanyBankAccount(transaction.CompanyBankAccountTo.Id)
	if err != nil {
		return entity.Transaction{}, errors.New("bank account must exist")
	}
	db.mx.Lock()
	defer db.mx.Unlock()
	db.Transactions = append(db.Transactions, transaction)
	return transaction, nil
}

func (db *Database) GetOneTransaction(id int) (entity.Transaction, error) {
	db.mx.Lock()
	defer db.mx.Unlock()
	for _, c := range db.Transactions {
		if c.Id == id {
			return c, nil
		}
	}
	return entity.Transaction{}, errors.New("Transaction not found")
}

func (db *Database) UpdateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, u := range db.Transactions {
		if u.Id == transaction.Id {
			transaction.SetUpdatedAtNow()
			db.Transactions[i] = transaction
			return transaction, nil
		}
	}
	return entity.Transaction{}, errors.New("transaction not found")
}

func (db *Database) RemoveTransaction(transaction entity.Transaction) error {
	db.mx.Lock()
	defer db.mx.Unlock()
	for i, c := range db.Transactions {
		if c.Id == transaction.Id {
			if len(db.Transactions) > 1 {
				db.Transactions = append(db.Transactions[0:i], db.Transactions[i+1:]...)
				return nil
			} else {
				db.Transactions = make([]entity.Transaction, 0)
				return nil
			}
		}
	}
	return errors.New("transaction not found")
}

func(db *Database) CommitTransaction (transaction entity.Transaction) (entity.Transaction, error) {

	var err error
	if transaction.Sum > transaction.CompanyBankAccountFrom.Sum {
		return transaction, errors.New("soruce bank account balance error")
	}

	db.mx.Lock()
	transaction.CompanyBankAccountFrom.Sum = transaction.CompanyBankAccountFrom.Sum - transaction.Sum
	transaction.CompanyBankAccountTo.Sum = transaction.CompanyBankAccountTo.Sum + transaction.Sum
	db.mx.Unlock()

	transaction.CompanyBankAccountFrom , err = db.UpdateCompanyBankAccount(transaction.CompanyBankAccountFrom)
	if err != nil {
		return transaction, errors.New("UpdateCompanyBankAccountFrom error")
	}
	transaction.CompanyBankAccountTo , err = db.UpdateCompanyBankAccount(transaction.CompanyBankAccountTo)
	if err != nil {
		return transaction, errors.New("UpdateCompanyBankAccountTo error")
	}

	db.mx.Lock()
	transaction.InitiatorCustomer.Transactions = append(transaction.InitiatorCustomer.Transactions, transaction)
	db.mx.Unlock()


	transaction.InitiatorCustomer, err = db.UpdateCustomer(transaction.InitiatorCustomer)
	if err != nil {
		return transaction, errors.New("UpdateInitiatorCustomer error")
	}

	transaction, err = db.UpdateTransaction(transaction)
	if err != nil {
		return transaction, errors.New("UpdateTransaction error")
	}
	return transaction, nil
}


