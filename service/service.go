package service

import (
	"v1/models"
	"v1/repositories"
)

type UserService struct {
	mysql *repositories.UserPostgresRepositories
}

func NewUserService(mysql *repositories.UserPostgresRepositories) *UserService {
	return &UserService{
		mysql: mysql,
	}
}
func (s *UserService) SetUser(user models.User) error {
	return s.mysql.InsertUser(user)
}
