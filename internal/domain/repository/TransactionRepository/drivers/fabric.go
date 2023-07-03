package drivers

import (
	cfg "repo/config"
	"repo/internal/DI"
	"repo/internal/domain/repository/TransactionRepository"
)

func Make(config *cfg.Configuration, di *DI.ServiceContainer) TransactionRepository.ITransactionRepository {
	switch config.Env {
	case cfg.ENV_TEST:
		config.Repository.Driver = cfg.DRIVER_MOCK
	}

	switch config.Repository.Driver {
	case cfg.DRIVER_MYSQL:
		return MakeMysqlDriver(config, di)
	case cfg.DRIVER_MOCK:
		return MakeMockDriver(config, di)
	default:
		return MakeMockDriver(config, di)
	}
}
