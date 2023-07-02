package svc

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/zeromicro/go-zero/core/logx"
	"gozero-asynq/asynqclient/internal/config"
	"net/http"
)

type ServiceContext struct {
	Config         config.Config
	AsynqClient    *asynq.Client
	AsynqScheduler *asynq.Scheduler
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Disable()

	redisOpt := asynq.RedisClientOpt{Addr: c.RedisConf.Host, Password: c.RedisConf.Pass}
	// 0. 启动dashboard
	startDashboard(redisOpt)

	// 1. 初始化client
	client := asynq.NewClient(redisOpt)
	// 2. 初始化scheduler
	scheduler := asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{
		LogLevel: asynq.InfoLevel,
		PostEnqueueFunc: func(taskInfo *asynq.TaskInfo, err error) {
			fmt.Printf("task id: %s, queue: %s, err: %+v\n", taskInfo.ID, taskInfo.Queue, err)
		},
	})

	return &ServiceContext{
		Config:         c,
		AsynqClient:    client,
		AsynqScheduler: scheduler,
	}
}

func startDashboard(r asynq.RedisConnOpt) {
	go func() {
		h := asynqmon.New(asynqmon.Options{
			RootPath:     "/dashboard", // RootPath specifies the root for asynqmon app
			RedisConnOpt: r,
			ReadOnly:     true,
		})

		// Note: We need the tailing slash when using net/http.ServeMux.
		http.Handle(h.RootPath()+"/", h)
		http.ListenAndServe(":8090", nil)
	}()
}
