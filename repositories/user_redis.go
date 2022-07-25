package repositories

import (
	"strconv"
	"time"
	"v1/dal"
	"v1/models"

	"github.com/go-redis/redis"
)

type UserRedisRepositories struct {
	databaseRedis *redis.Client
}

func NewUserRedisRepositories() *UserRedisRepositories {
	return &UserRedisRepositories{
		databaseRedis: dal.NewRedisClient(),
	}
}

func (r *UserRedisRepositories) InsertUser(user models.User) error {
	err := r.databaseRedis.Set(strconv.Itoa(user.Id), user, 12*time.Hour).Err()
	return err
}