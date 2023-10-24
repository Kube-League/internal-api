package sql

import (
	"gorm.io/gorm"
	"internal-api/src/config"
)

var DB *gorm.DB

func Setup() {
	DB = config.Cfg.Sql.Connect()
}
