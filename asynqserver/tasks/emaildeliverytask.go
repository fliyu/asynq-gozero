package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

// EmailDeliveryPayload 立即发送邮件负载
type EmailDeliveryPayload struct {
	Email   string
	Title   string
	Content string
}

// HandleEmailDeliveryTask 处理发送邮件任务
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var data EmailDeliveryPayload
	err := json.Unmarshal(t.Payload(), &data)
	if err != nil {
		return err
	}

	// do something
	fmt.Printf("time: %s,email delivery data: %+v\n", time.Now().Format("2006-01-02 15:04:05"), data)
	return nil
}
