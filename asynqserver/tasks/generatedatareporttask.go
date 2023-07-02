package tasks

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

// HandleGenerateDataReportTask 处理生成报表任务
func HandleGenerateDataReportTask(ctx context.Context, t *asynq.Task) error {
	// do something
	fmt.Printf("generate data report, time: %s\n", time.Now().Format(time.RFC3339))
	time.Sleep(time.Second * 30)
	return nil
}
