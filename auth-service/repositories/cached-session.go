package repositories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"auth-service/models"

	"github.com/redis/go-redis/v9"
)

type cachedSessionRepository struct {
	next  SessionRepository
	redis *redis.Client
}

func NewCachedSessionRepository(
	next SessionRepository,
	redis *redis.Client,
) SessionRepository {

	return &cachedSessionRepository{
		next:  next,
		redis: redis,
	}
}

func (r *cachedSessionRepository) FindByTokenHash(
	ctx context.Context,
	hash string,
) (*models.Session, error) {

	key := "session:" + hash

	data, err := r.redis.Get(ctx, key).Bytes()

	if err == nil {
		log.Printf("session was already available in cache with value %s", string(data))
		var session models.Session

		if err := json.Unmarshal(data, &session); err == nil {
			return &session, nil
		}
	}

	log.Print("session was NOT available in cache, retrieving persisted value...")
	session, err := r.next.FindByTokenHash(ctx, hash)

	if err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(session)

	ttl := time.Until(session.ExpiresAt)

	r.redis.Set(
		ctx,
		key,
		bytes,
		ttl,
	)

	return session, nil
}

func (r *cachedSessionRepository) Create(
	ctx context.Context,
	session *models.Session,
) error {

	err := r.next.Create(ctx, session)

	if err != nil {
		return err
	}

	key := "session:" + session.SessionTokenHash

	bytes, _ := json.Marshal(session)

	ttl := time.Until(session.ExpiresAt)

	r.redis.Set(
		ctx,
		key,
		bytes,
		ttl,
	)

	return nil
}

func (r *cachedSessionRepository) Revoke(
	ctx context.Context,
	hash string,
) error {

	err := r.next.Revoke(ctx, hash)

	if err != nil {
		return err
	}

	return r.redis.Del(
		ctx,
		"session:"+hash,
	).Err()
}
