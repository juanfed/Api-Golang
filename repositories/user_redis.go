package repositories

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	usuario := &user
	usuario.Id = user.Id

	userData, _ := json.Marshal(usuario)

	fmt.Println("Dato que le llega al redis: ", user)
	err := r.databaseRedis.Set(strconv.Itoa(usuario.Id), userData, 0).Err()
	return err
}
