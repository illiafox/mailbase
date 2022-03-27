package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/illiafox/mailbase/shared/public"
	"strconv"
	"time"
)

type Forgot struct {
	Client *redis.Client
	Expire time.Duration
}

func (f *Forgot) New(userid int, key string) error {
	key, err := EventJson(ForgotPass, key)
	if err != nil {
		return err
	}
	return f.Client.SetEX(context.Background(), key, userid, f.Expire).Err()
}

func (f *Forgot) Get(key string) (int, error) {
	key, err := EventJson(ForgotPass, key)
	if err != nil {
		return -1, err
	}
	n, err := f.Client.GetDel(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return -1, public.Register.KeyNotFound
		} else {
			return -1, public.NewInternalWithError(err)
		}
	}
	return strconv.Atoi(n)
}
