package tasks

import (
	"fmt"
	"github.com/hibiken/asynq"
	"gozero-asynq/types"
	"time"
)

// NewGenerateDataReportTask 创建一个定时生成报表任务，处理所有用户
func NewGenerateDataReportTask() (*asynq.Task, error) {
	return asynq.NewTask(types.TypeGenerateDataReport, nil), nil
}

// RegisterGenerateDataReportTask 注册生成报表任务
func RegisterGenerateDataReportTask(scheduler *asynq.Scheduler, cronspec string) error {
	task, err := NewGenerateDataReportTask()
	if err != nil {
		return err
	}

	entryId, err := scheduler.Register(cronspec, task, asynq.MaxRetry(3))
	if err != nil {
		return err
	}
	fmt.Printf("register entryId: %s, time: %s\n", entryId, time.Now().Format("2006-01-02 15:04:05"))

	return nil
}
