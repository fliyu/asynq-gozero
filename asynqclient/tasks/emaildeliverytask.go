package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"gozero-asynq/types"
)

// EmailDeliveryPayload 立即发送邮件负载
type EmailDeliveryPayload struct {
	Email   string
	Title   string
	Content string
}

// NewEmailDeliveryTask 创建一个立即发送邮件任务，为指定用户发送
func NewEmailDeliveryTask(data EmailDeliveryPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// asynq.NewTask
	// 传入任务名称、序列化的负载信息以及opts配置项
	return asynq.NewTask(types.TypeEmailDelivery, payload), nil
}

// RegisterEmailDeliveryTask 立即发送邮件任务
func RegisterEmailDeliveryTask(client *asynq.Client, data EmailDeliveryPayload) error {
	task, err := NewEmailDeliveryTask(data)
	if err != nil {
		return err
	}

	// Enqueue task to be processed immediately.
	taskInfo, err := client.Enqueue(task, asynq.MaxRetry(1))
	if err != nil {
		return err
	}
	fmt.Printf("enqueued task id: %s, queue: %s\n", taskInfo.ID, taskInfo.Queue)

	return nil
}
