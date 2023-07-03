package mysqlStorage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"repo/config"
	"repo/internal/DI"
	"time"
)

func MakeMysqlDb(cfg *config.Configuration, di *DI.ServiceContainer) *sql.DB {
	mysqlConfig := mysql.Config{
		User:                 cfg.Database.Mysql.Username,
		Passwd:               cfg.Database.Mysql.Password,
		Net:                  "tcp",
		Addr:                 cfg.Database.Mysql.Host,
		DBName:               cfg.Database.Mysql.Schema,
		Timeout:              cfg.Database.Mysql.DailTimeout,
		ReadTimeout:          cfg.Database.Mysql.ReadTimeout,
		WriteTimeout:         cfg.Database.Mysql.WriteTimeout,
		InterpolateParams:    true,
		MultiStatements:      true,
		ParseTime:            false,
		AllowNativePasswords: true,
	}

	db, err := sql.Open(cfg.Database.DriverName, mysqlConfig.FormatDSN())
	if err != nil {
		fmt.Print(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(),5* time.Second)
	defer cancel()

	var t time.Time

	result := db.QueryRowContext(ctx,"select now()")
	result.Scan(&t)
	fmt.Println(t)

	return db
}
