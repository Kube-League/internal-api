package main

import (
	"internal-api/src/config"
	"internal-api/src/db/sql"
)

func main() {
	config.Setup()
	sql.Setup()

	err := sql.DB.AutoMigrate(&sql.Match{}) // it will also set up team, player, etc.
	if err != nil {
		panic(err)
	}
}
