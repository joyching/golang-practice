package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joyching/golang-practice/config"
)

func Connect() (*sql.DB, error) {
	config := config.GetConfig()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.GetString("db.username"),
		config.GetString("db.password"),
		config.GetString("db.host"),
		config.GetString("db.port"),
		config.GetString("db.database"),
	)

	return sql.Open("mysql", dataSourceName)
}
