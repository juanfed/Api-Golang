package dal

import (
	"fmt"
	"v1/constants"
)

<<<<<<< HEAD
// paso todo los datos necesarios para la coneccion con la base de datos
=======
>>>>>>> aa586aab2b9c4c9402eb8cb8ac81b5164c597c6b
func GetSqlConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		constants.DBHost,
		constants.DBPort,
		constants.DBUser,
		constants.DBName,
		constants.DBPassword)
	return dataBase
}
