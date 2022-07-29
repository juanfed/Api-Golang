package service

import (
	"v1/models"
	"v1/repositories"
)

type UserService struct {
	mysql *repositories.UserMysqlRepositories
	redis *repositories.UserRedisRepositories
}

func NewUserService(mysql *repositories.UserMysqlRepositories, redis *repositories.UserRedisRepositories) *UserService {
	return &UserService{
		mysql: mysql,
		redis: redis,
	}
}
func NewUserServiceRedis(redis *repositories.UserRedisRepositories) *UserService {
	return &UserService{
		redis: redis,
	}
}

func (s *UserService) Set(user models.User) error {
	err := s.mysql.Set(user)
	if err != nil {
		return err
	}

	return s.redis.Set(user)
}

func (s *UserService) Get(id int) (models.User, error) {
	// primero debo deconsiltar si esta ese dato en el redis y en caso de que no este debo de ir a consultar entonces en la base de datos que es mi fuente de verdad

	// espacio para la consutaen el redis

	user, err := s.mysql.Get(id)
	if err != nil {
		return models.User{}, err
	}

	// en caso de que me toca

	return user, nil
}

// Cuando se hace un update en el postgres, se debe de borrar el dato en el redis para evitar problemas por condicion de carrera
func (s *UserService) Update(user models.User) error {
	err := s.mysql.Update(user)
	if err != nil {
		return err
	}

	return s.redis.Delete(user.Id)
}

func (s *UserService) Delete(id int) error {
	err := s.mysql.Delete(id)
	if err != nil {
		return err
	}

	return s.mysql.Delete(id)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.mysql.GetAll()
}
