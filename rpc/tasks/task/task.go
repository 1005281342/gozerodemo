package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/1005281342/gozerodemo/rpc/tasks/tasks"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

const (
	RedisAddr   = "192.168.64.1:6379"
	Concurrency = 16
)

type PushPayload struct {
	User      string `json:"user"`
	Message   string `json:"message"`
	At        int64  `json:"at"`
	In        int64  `json:"in"`
	Timestamp int64  `json:"timestamp"`
}

func NewPushTask(payload string) (*asynq.Task, error) {
	var p = PushPayload{}
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		return nil, err
	}
	var opt []asynq.Option
	if p.In != 0 {
		opt = append(opt, asynq.ProcessAt(time.Now().Add(time.Duration(p.In)*time.Millisecond)))
	} else if p.At != 0 {
		opt = append(opt, asynq.ProcessAt(time.Unix(p.At, 0)))
	}
	opt = append(opt, asynq.Queue(tasks.TaskType_Push.String()))

	return asynq.NewTask(tasks.TaskType_Push.String(), []byte(payload), opt...), nil
}

func (p PushPayload) HandlePushTask(ctx context.Context, t *asynq.Task) error {
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Pushing: user=%s, message=%s", p.User, p.Message)
	return nil
}
