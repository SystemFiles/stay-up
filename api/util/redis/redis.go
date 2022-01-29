package rclient

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/systemfiles/stay-up/api/config"
	"github.com/systemfiles/stay-up/api/models"
)

var ctx = context.Background()

var (
	instance *redis.Client
	once sync.Once
)

func getClient() (*redis.Client, error) {
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", config.App.RDBHost, config.App.RDBPort),
			Password: config.App.RDBPass,
			DB: 0,
		})

		instance = rdb
	})

	return instance, nil
}

func GetAll(dest *[]models.Service) error {
	c, err := getClient()
	if err != nil {
		return err
	}

	keys, err := c.Keys(ctx, "*").Result()
	if err != nil {
		return err
	}

	var sList []models.Service
	for _, key := range(keys) {
		var s models.Service
		if err := Get(key, &s); err != nil {
			return err
		}
		
		sList = append(sList, s)
	}

	*dest = sList
	return nil
}

func Set(key string, value interface{}) error {
	c, err := getClient()
	if err != nil {
		return err
	}

	value_json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Set(ctx, key, value_json, 0).Err()
}

func Get(key string, dest interface{}) error {
	c, err := getClient()
	if err != nil {
		return err
	}

	value, err := c.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	value_b := []byte(value)
	return json.Unmarshal(value_b, dest)
}

func Delete(key string) error {
	c, err := getClient()
	if err != nil {
		return err
	}

	return c.Del(ctx, key).Err()
}