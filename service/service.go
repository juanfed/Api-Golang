package service

import (
	"v1/models"
	"v1/repositories"
)

type UserService struct {
	mysql *repositories.UserPostgresRepositories
	redis *repositories.UserRedisRepositories
}

func NewUserService(mysql *repositories.UserPostgresRepositories) *UserService {
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
