package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"time"
)

type Verify struct {
	Client *redis.Client
	Expire time.Duration
}

func (v *Verify) New(user model.Users, key string) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key, err = EventJSON(VerifyUser, key)
	if err != nil {
		return err
	}
	return v.Client.SetEX(context.Background(), key, data, v.Expire).Err()
}

func (v *Verify) Get(key string) (model.Users, error) {
	var user model.Users

	key, err := EventJSON(VerifyUser, key)
	if err != nil {
		return user, public.NewInternalWithError(err)
	}

	data, err := v.Client.GetDel(context.Background(), key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return user, public.Register.KeyNotFound
		}
		return user, public.NewInternalWithError(err)
	}

	err = json.Unmarshal(data, &user)
	return user, err
}
