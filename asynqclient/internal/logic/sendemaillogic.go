package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"gozero-asynq/asynqclient/tasks"
	"time"

	"gozero-asynq/asynqclient/internal/svc"
	"gozero-asynq/asynqclient/internal/types"
	common "gozero-asynq/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SendEmail 发送邮件
func (l *SendEmailLogic) SendEmail(req *types.SendEmailRequest) (resp *types.SendEmailResponse, err error) {
	if req.IsDelay {
		// 模拟定时发送
		data := tasks.EmailDeliveryPayload{
			Email:   req.Email,
			Title:   req.Title,
			Content: req.Content,
		}
		payload, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		task := asynq.NewTask(common.TypeEmailDelivery, payload)
		processIn := asynq.ProcessAt(time.Unix(req.Time, 0))

		enqueue, err := l.svcCtx.AsynqClient.EnqueueContext(l.ctx, task, processIn)
		if err != nil {
			return nil, err
		}
		l.Infof("enqueue: %v", enqueue.ID)
	}
	l.Infof("send email success, time:%s", time.Now().Format("2006-01-02 15:04:05"))
	return
}
