package dal

import (
	"database/sql"
	"fmt"
	"log"
	"v1/constants"
)

func GetConnetMySql() string {
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		constants.DBUser,
		constants.DBPassword,
		constants.DBHost,
		constants.DBPort,
		constants.DBName,
	)

	return database
}
func NewDatabaseSql() *sql.DB {
	db, err := sql.Open(constants.DBType, GetConnetMySql())
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connect to database")

	return db
}
