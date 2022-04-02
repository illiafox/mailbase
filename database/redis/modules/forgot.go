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

type Forgot struct {
	Client *redis.Client
	Expire time.Duration
}

// New forgot password event
// key: hash key
func (f *Forgot) New(userID int, key string) error {
	key, err := event.JSON(event.ForgotPass, key)
	if err != nil {
		return err
	}
	return public.NewInternalWithError(
		f.Client.SetEX(context.Background(), key, userID, f.Expire).Err(),
	)
}

// Get returns user id
// key: hash key
func (f *Forgot) Get(key string) (int, error) {
	key, err := event.JSON(event.ForgotPass, key)
	if err != nil {
		return -1, err
	}
	n, err := f.Client.GetDel(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return -1, public.Register.KeyNotFound
		}
		return -1, public.NewInternalWithError(err)
	}
	return strconv.Atoi(n)
}
