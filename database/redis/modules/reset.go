package modules

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/illiafox/mailbase/database/redis/event"
	"github.com/illiafox/mailbase/shared/public"
	"strconv"
	"time"
)

type Reset struct {
	Client *redis.Client
	Expire time.Duration
}

func (r *Reset) New(userid int, key string) error {
	key, err := event.EventJSON(event.ResetPass, key)
	if err != nil {
		return err
	}
	return r.Client.SetEX(context.Background(), key, userid, r.Expire).Err()
}

func (r *Reset) Get(key string) (int, error) {
	key, err := event.EventJSON(event.ResetPass, key)
	if err != nil {
		return -1, err
	}

	n, err := r.Client.GetDel(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return -1, public.Register.KeyNotFound
		}
		return -1, public.NewInternalWithError(err)
	}

	return strconv.Atoi(n)
}
