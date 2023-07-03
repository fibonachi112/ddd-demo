package fabric

import (
	"fmt"
	"repo/config"
	"repo/internal/DI"
	companyAccountRepositoryFabric "repo/internal/domain/repository/CompanyAccountRepository/drivers"
	companyBankAccountRepositoryFabric "repo/internal/domain/repository/CompanyBankAccountRepository/drivers"
	companyRepositoryFabric "repo/internal/domain/repository/CompanyRepository/drivers"
	customerRepositoryFabric "repo/internal/domain/repository/CustomerRepository/drivers"
	transactionRepositoryFabric "repo/internal/domain/repository/TransactionRepository/drivers"
	userRepositoryFabric "repo/internal/domain/repository/UserRepository/drivers"
	"repo/internal/pkg/inmemoryStorage"
	"repo/internal/pkg/mysqlStorage"
	"sync"
)

var OnceContainer sync.Once
var Container DI.ServiceContainer

func Make() DI.ServiceContainer {

	OnceContainer.Do(func() {
		fmt.Print("creating container")
		Container = DI.ServiceContainer{}
		Container.MysqlDatabase = mysqlStorage.MakeMysqlDb(config.Config, &Container)
		Container.InMemoryStorage = inmemoryStorage.MakeDatabase()
		Container.UserRepository = userRepositoryFabric.Make(config.Config, &Container)
		Container.CompanyAccountRepository = companyAccountRepositoryFabric.Make(config.Config, &Container)
		Container.CompanyRepository = companyRepositoryFabric.Make(config.Config, &Container)
		Container.CustomerRepository = customerRepositoryFabric.Make(config.Config, &Container)
		Container.CompanyBankAccountRepository = companyBankAccountRepositoryFabric.Make(config.Config, &Container)
		Container.TransactionRepository = transactionRepositoryFabric.Make(config.Config, &Container)
	})

	return Container
}
