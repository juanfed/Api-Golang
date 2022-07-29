package repositories

import (
	"database/sql"
	"fmt"
	"v1/dal"
	"v1/models"

	// conexion del driver
	_ "github.com/go-sql-driver/mysql"
)

type UserMysqlRepositories struct {
	database *sql.DB
}

func NewUserMysqlRepositories() *UserMysqlRepositories {
	return &UserMysqlRepositories{
		database: dal.NewDatabaseSql(),
	}
}
func (sq *UserMysqlRepositories) Set(user models.User) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(
			`insert into user (id, name, last_name) values(%d, "%s", "%s")`,
			user.Id,
			user.Name,
			user.LastName,
		),
	)

	return err
}
func (sq *UserMysqlRepositories) Delete(id int) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(
			`delete from user where id='%d'`,
			id,
		),
	)
	return err
}
