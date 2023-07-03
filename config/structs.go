package config

import "time"

type Mysql struct {
	Username     string
	Password     string
	Host         string
	Schema       string
	DailTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// конфиг базы данных
type Database struct {
	DriverName string
	Mysql *Mysql
}

// конфиг http сервера
type HttpApp struct {
	ListenAddr string
	ListenPort string
	AppName    string
}

type Repository struct {
	Driver string
}

// весь конфиг
type Configuration struct {
	Env        string
	Database   *Database
	HttpApp    *HttpApp
	Repository *Repository
}
