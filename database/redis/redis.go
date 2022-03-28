package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/illiafox/mailbase/database/redis/modules"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/util/config"
	"time"
)

type Redis struct {
	Client *redis.Client

	Forgot modules.Forgot
	Reset  modules.Reset
	Verify modules.Verify
}

func NewRedis(client *redis.Client, conf config.Config) Redis {
	expire := time.Duration(public.Redis.ExpireSeconds) * time.Second

	return Redis{
		Client: client,

		Forgot: modules.Forgot{
			Client: client,
			Expire: expire,
		},

		Reset: modules.Reset{
			Client: client,
			Expire: expire,
		},

		Verify: modules.Verify{
			Client: client,
			Expire: expire,
		},
	}
}
