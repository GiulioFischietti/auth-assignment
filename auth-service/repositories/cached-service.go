package repositories

import (
	"context"
	"log"

	"auth-service/models"

	"github.com/redis/go-redis/v9"
)

type cachedServiceRepository struct {
	next  ServiceRepository
	redis *redis.Client
}

func NewCachedServiceRepository(
	next ServiceRepository,
	redis *redis.Client,
) ServiceRepository {

	return &cachedServiceRepository{
		next:  next,
		redis: redis,
	}
}

func (r *cachedServiceRepository) FindByName(
	ctx context.Context,
	name string,
) (*models.Service, error) {

	key := "service:" + name

	data, err := r.redis.Get(
		ctx,
		key,
	).Bool()

	if err == nil {
		log.Printf("cached service found in cache with value %s", data)
		service := &models.Service{
			Name:   name,
			Active: data,
		}
		return service, nil
	}

	log.Print("cached service NOT found in cache! retrieving it from postgres...")
	service, err := r.next.FindByName(
		ctx,
		name,
	)

	if err != nil {
		return nil, err
	}

	r.redis.Set(
		ctx,
		key,
		true,
		0,
	)

	return service, nil
}
