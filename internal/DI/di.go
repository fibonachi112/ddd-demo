package DI

import (
	"database/sql"
	"repo/internal/domain/repository/CompanyAccountRepository"
	"repo/internal/domain/repository/CompanyBankAccountRepository"
	"repo/internal/domain/repository/CompanyRepository"
	"repo/internal/domain/repository/CustomerRepository"
	"repo/internal/domain/repository/TransactionRepository"
	"repo/internal/domain/repository/UserRepository"
	"repo/internal/pkg/inmemoryStorage"
)

type ServiceContainer struct {
	MysqlDatabase                *sql.DB
	InMemoryStorage              *inmemoryStorage.Database
	UserRepository               UserRepository.IUserRepository
	CompanyAccountRepository     CompanyAccountRepository.ICompanyAccountRepository
	CompanyRepository            CompanyRepository.ICompanyRepository
	CustomerRepository           CustomerRepository.ICustomerRepository
	CompanyBankAccountRepository CompanyBankAccountRepository.ICompanyBankAccountRepository
	TransactionRepository        TransactionRepository.ITransactionRepository
}
