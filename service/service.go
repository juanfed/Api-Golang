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

func (s *UserService) SetUser(user models.User) error {
	err := s.mysql.InsertUser(user)
	if err != nil {
		return err
	}
	return s.redis.InsertUser(user)
}
