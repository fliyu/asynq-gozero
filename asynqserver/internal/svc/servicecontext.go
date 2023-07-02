package svc

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"gozero-asynq/asynqserver/internal/config"
	"gozero-asynq/asynqserver/tasks"
	common "gozero-asynq/types"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Disable()

	redisOpt := asynq.RedisClientOpt{Addr: c.RedisConf.Host, Password: c.RedisConf.Pass}
	// 1. asynq 配置
	conf := asynq.Config{
		// Specify how many concurrent workers to use
		Concurrency: 10,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			fmt.Printf("task type: %s, payload: %s, err: %+v\n", task.Type(), string(task.Payload()), err)
		}),
	}

	// 2. 注册全局任务
	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	RegisterTaskHandlers(mux)

	// 3. 启动 asynq 服务
	server := asynq.NewServer(redisOpt, conf)
	go func() {
		if err := server.Run(mux); err != nil {
			panic(err)
		}
	}()

	return &ServiceContext{
		Config: c,
	}
}

// RegisterTaskHandlers 注册任务
func RegisterTaskHandlers(mux *asynq.ServeMux) {
	mux.HandleFunc(common.TypeEmailDelivery, tasks.HandleEmailDeliveryTask)
	mux.HandleFunc(common.TypeGenerateDataReport, tasks.HandleGenerateDataReportTask)
}
