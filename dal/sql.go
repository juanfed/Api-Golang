package dal

import (
	"fmt"
	"v1/constants"
)

// paso todo los datos necesarios para la coneccion con la base de datos
func GetSqlConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		constants.DBHost,
		constants.DBPort,
		constants.DBUser,
		constants.DBName,
		constants.DBPassword)
	return dataBase
}
