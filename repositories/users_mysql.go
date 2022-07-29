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

func NewUserMysqlRepositories(mysql *sql.DB) *UserMysqlRepositories {
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

func (sq *UserMysqlRepositories) Update(user models.User) error {
	_, err := sq.database.Exec(
		fmt.Sprintf(
			`update user set name='%s', last_name='%s' where id=%d`,
			user.Name,
			user.LastName,
			user.Id,
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
