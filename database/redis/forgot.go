package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mailbase/shared/public"
	"strconv"
)

func (r *Redis) NewForgotPass(userid int, key string) error {
	key, err := EventJson(ForgotPass, key)
	if err != nil {
		return err
	}
	return r.Client.SetEX(context.Background(), key, userid, r.Expire).Err()
}

func (r *Redis) GetForgotPass(key string) (int, error) {
	key, err := EventJson(ForgotPass, key)
	if err != nil {
		return -1, err
	}
	n, err := r.Client.GetDel(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return -1, public.Register.KeyNotFound
		} else {
			return -1, public.NewInternalWithError(err)
		}
	}

	return strconv.Atoi(n)
}
