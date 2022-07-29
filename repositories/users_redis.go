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
	database *redis.Client
}

func NewUserRedisRepositories() *UserRedisRepositories {
	return &UserRedisRepositories{
		database: dal.NewRedisClient(),
	}
}

func (r *UserRedisRepositories) Set(user models.User) error {
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.database.Set(fmt.Sprintf("user:%d", user.Id), value, 12*time.Hour).Err()
}

func (r *UserRedisRepositories) Delete(id int) error {
	return r.database.Del(fmt.Sprintf("user:%d", id)).Err()
}
