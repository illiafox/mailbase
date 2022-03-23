package redis

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	Client *redis.Client
	Expire time.Duration // Time for buf expiration IN SECONDS
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

func EventJson(task Task, key string) (string, error) {
	data, err := json.Marshal(Event{task, key})
	if err != nil {
		return "", err
	}
	return string(data), nil
}
