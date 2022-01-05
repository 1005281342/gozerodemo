package task

import (
	"context"
	"fmt"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"
	"github.com/hibiken/asynq"
	"log"
)

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case tasks.TaskType_Push.String():
		var p PushPayload
		return p.HandlePushTask(ctx, t)
	}
	return fmt.Errorf("unexpected task type: %s", t.Type())
}

func registerWorker() {
	var srv = asynq.NewServer(
		asynq.RedisClientOpt{Addr: RedisAddr},
		asynq.Config{Concurrency: Concurrency, Queues: map[string]int{
			tasks.TaskType_Push.String(): 6,
		}},
	)

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}

// DeleteTask 从队列中移除任务
func DeleteTask(qname string, uuid string) error {
	var ins = asynq.NewInspector(asynq.RedisClientOpt{Addr: RedisAddr})
	return ins.DeleteTask(qname, uuid)
}
