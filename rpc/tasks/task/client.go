package task

import "github.com/hibiken/asynq"

var client *asynq.Client

func initClient() {
	client = asynq.NewClient(asynq.RedisClientOpt{Addr: RedisAddr})
}

func GetClient() *asynq.Client {
	if client == nil {
		initClient()
	}

	return client
}
