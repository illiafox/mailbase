package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"mailbase/database/mysql/model"
	"mailbase/shared/public"
	"time"
)

type Redis struct {
	Client *redis.Client
	Expire time.Duration // Time for buf expiration IN SECONDS
}

func (r *Redis) NewBuf(user model.Users, key string) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.Client.SetEX(context.Background(), key, data, r.Expire).Err()
}

func (r *Redis) GetBuf(key string) (model.Users, error) {
	var user model.Users

	data, err := r.Client.GetDel(context.Background(), key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return user, public.Register.KeyNotFound
		} else {
			return user, public.NewInternalWithError(err)
		}
	}

	err = json.Unmarshal(data, &user)
	return user, err
}
