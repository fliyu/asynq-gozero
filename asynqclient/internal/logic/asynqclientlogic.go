package logic

import (
	"context"

	"gozero-asynq/asynqclient/internal/svc"
	"gozero-asynq/asynqclient/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AsynqclientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqclientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqclientLogic {
	return &AsynqclientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsynqclientLogic) Asynqclient(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
