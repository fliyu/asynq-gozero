package main

import (
	"flag"
	"fmt"
	"github.com/hibiken/asynq"
	"gozero-asynq/asynqclient/tasks"

	"gozero-asynq/asynqclient/internal/config"
	"gozero-asynq/asynqclient/internal/handler"
	"gozero-asynq/asynqclient/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/asynqclient-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册全局性任务
	RegisterGlobalTasks(ctx.AsynqClient, ctx.AsynqScheduler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// RegisterGlobalTasks 注册全局性任务
func RegisterGlobalTasks(client *asynq.Client, scheduler *asynq.Scheduler) {
	// 1. 注册立即执行的任务
	err := tasks.RegisterEmailDeliveryTask(client, tasks.EmailDeliveryPayload{
		Email:   "test@qq.com",
		Title:   "定时发送",
		Content: "Hello World",
	})
	if err != nil {
		return
	}

	// 2. 注册周期性任务
	cronspec := "* * * * *" //"1 2 * * *" // 每天 02:01:00 执行定时任务，最大重试次数3次.
	err = tasks.RegisterGenerateDataReportTask(scheduler, cronspec)
	if err != nil {
		return
	}

	go func() {
		err = scheduler.Run() // 这里需要调用run方法，会进行阻塞，别忘记了
		if err != nil {
			fmt.Printf("scheduler run error: %v", err)
			return
		}
		defer scheduler.Shutdown()
	}()
}
