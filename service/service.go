package service

import (
	"v1/models"
	"v1/repositories"
)

type UserService struct {
	mysql *repositories.UserMysqlRepositories
	redis *repositories.UserRedisRepositories
}

func NewUserService(mysql *repositories.UserMysqlRepositories) *UserService {
	return &UserService{
		mysql: mysql,
	}
}
func NewUserServiceRedis(redis *repositories.UserRedisRepositories) *UserService {
	return &UserService{
		redis: redis,
	}
}

func (s *UserService) SetUser(user models.User) error {
	return s.mysql.InsertUser(user)
}
func (r *UserService) SetUserRedis(user models.User) error {
	return r.redis.InsertUser(user)
}
