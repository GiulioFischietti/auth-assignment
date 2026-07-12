package repositories

import (
	"context"
	"encoding/json"
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
	).Bytes()

	if err == nil {
		log.Printf("cached service found in cache with value %s", string(data))
		service := &models.Service{}

		if err := json.Unmarshal(
			data,
			service,
		); err == nil {

			return service, nil
		}
	}

	log.Print("cached service NOT found in cache! retrieving it from postgres...")
	service, err := r.next.FindByName(
		ctx,
		name,
	)

	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(service)

	if err == nil {
		r.redis.Set(
			ctx,
			key,
			bytes,
			0,
		)
	}

	return service, nil
}
