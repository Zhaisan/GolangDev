package redis_cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"lectures-6/internal/cache"
	"lectures-6/internal/models"
	"time"
)

func (rc RedisCache) Laptops() cache.LaptopsCacheRepo {
	if rc.laptops == nil {
		rc.laptops = newLaptopsRepo(rc.client, rc.expires)
	}

	return rc.laptops
}

type LaptopsRepo struct {
	client  *redis.Client
	expires time.Duration
}

func newLaptopsRepo(client *redis.Client, exp time.Duration) cache.LaptopsCacheRepo {
	return &LaptopsRepo{
		client:  client,
		expires: exp,
	}
}

func (lr LaptopsRepo) Set(ctx context.Context, key string, value []*models.Laptop) error {
	laptopBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = lr.client.Set(ctx, key, laptopBytes, lr.expires*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (lr LaptopsRepo) Get(ctx context.Context, key string) ([]*models.Laptop, error) {
	result, err := lr.client.Get(ctx, key).Result()
	switch err {
	case nil:
		break
	case redis.Nil:
		return nil, nil
	default:
		return nil, err
	}

	laptops := make([]*models.Laptop, 0)
	if err = json.Unmarshal([]byte(result), &laptops); err != nil {
		return nil, err
	}

	return laptops, nil
}
