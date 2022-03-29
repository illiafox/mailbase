package event

import "encoding/json"

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

func JSON(task Task, key string) (string, error) {
	data, err := json.Marshal(Event{task, key})
	if err != nil {
		return "", err
	}
	return string(data), nil
}
