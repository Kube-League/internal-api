package config

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	Cfg        *Config
	defaultCfg = Config{
		Sql: SqlCredentials{
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "admin",
			DbName:   "InternalAPI",
		},
	}
)

type Config struct {
	Sql SqlCredentials
}

type SqlCredentials struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
}

func (c *SqlCredentials) Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%u sslmode=disable",
		c.Host, c.User, c.Password, c.Port, c.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Setup() {
	data, err := os.ReadFile("/data/config.toml")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
		err = os.Mkdir("/data", 0744)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
		data, err = toml.Marshal(&defaultCfg)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("/data/config.toml", data, 0755)
		if err != nil {
			panic(err)
		}
	}
	err = toml.Unmarshal(data, &Cfg)
	if err != nil {
		panic(err)
	}
}
