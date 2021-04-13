package storages

import (
	"fmt"

	"boiler_plate.com/src/configuration/environments"
	"github.com/go-redis/redis"
)

type IRedisRepository interface {
	Save(key, value string) error
	Get(key string) (string, error)
}

type redisRepository struct {
	client *redis.Client
}

func (rr *redisRepository) Save(key, value string) error {
	err := rr.client.Set(key, value, 0).Err()
	return err
}

func (rr *redisRepository) Get(key string) (string, error) {
	result, err := rr.client.Get(key).Result()

	return result, err
}

func NewRedisRepository() IRedisRepository {
	env := environments.GetEnvironment()
	redisAddr := fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return &redisRepository{
		client: client,
	}
}
