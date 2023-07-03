package config

import (
	"fmt"
	"sync"
)

var Config *Configuration
var OnceConfig sync.Once

const ENV_PROD = "prod"
const ENV_DEV = "dev"
const ENV_TEST = "test"

const DRIVER_MYSQL = "mysql"
const DRIVER_MOCK = "mock"

func init() {
	OnceConfig.Do(func() {
		fmt.Println("config loaded")
		Config = &Configuration{
			Env: ENV_DEV,
			Database: &Database{
				DriverName: "mysql",
				Mysql: &Mysql{
					Username:     "",
					Password:     "",
					Host:         "host:port",
					Schema:       "databasename",
					DailTimeout:  0,
					ReadTimeout:  0,
					WriteTimeout: 0,
				},
			},

			Repository: &Repository{
				Driver: "mysql",
			},
			HttpApp: &HttpApp{
				ListenAddr: "0.0.0.0",
				ListenPort: ":8091",
				AppName:    "RepositoryApp",
			},
		}
	})
}
