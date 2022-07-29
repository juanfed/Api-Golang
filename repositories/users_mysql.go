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

func (sq *UserMysqlRepositories) Get(id int) (models.User, error) {
	value, err := sq.database.Query(
		fmt.Sprintf(
			`select id, name, last_name from user where id=%d`,
			id,
		),
	)
	if err != nil {
		return models.User{}, err
	}

	// instancio la variable que será un model de la struct user para pasarle los datos que capturo del mysql
	user := models.User{}
	if value.Next() {
		// paso los valores a la variable user que instancie de la struct de user
		err = value.Scan(&user.Id, &user.Name, &user.LastName)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, err
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

func (sq *UserMysqlRepositories) GetAll() ([]models.User, error) {
	rows, err := sq.database.Query("select * from user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	user := models.User{}
	data := []models.User{}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.LastName)
		if err != nil {
			return nil, err
		}
		data = append(data, user)
	}
	fmt.Println(data)

	return data, err
}
