package repositories

import (
	"encoding/json"
	"fmt"
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

	value, err := json.Marshal(user)

	if err != nil {
		return err
	}

	return r.databaseRedis.Set(fmt.Sprintf("user:%d", user.Id), value, 12*time.Hour).Err()
}
