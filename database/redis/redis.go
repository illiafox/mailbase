package redis

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/util/config"
	"time"
)

type Redis struct {
	Client *redis.Client

	Forgot Forgot
	Reset  Reset
	Verify Verify
}

func NewRedis(client *redis.Client, conf config.Config) Redis {
	expire := time.Duration(public.Redis.ExpireSeconds) * time.Second

	return Redis{
		Client: client,

		Forgot: Forgot{
			Client: client,
			Expire: expire,
		},

		Reset: Reset{
			Client: client,
			Expire: expire,
		},

		Verify: Verify{
			Client: client,
			Expire: expire,
		},
	}
}

type Event struct {
	Task Task
	Key  string
}

type Task string

const (
	VerifyUser = Task("verify_user")
	ForgotPass = Task("forgot_password")
	ResetPass  = Task("reset_password")
)

func EventJSON(task Task, key string) (string, error) {
	data, err := json.Marshal(Event{task, key})
	if err != nil {
		return "", err
	}
	return string(data), nil
}
