package logic

import (
	"context"

	"gozero-asynq/asynqserver/asynqserver"
	"gozero-asynq/asynqserver/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *asynqserver.Request) (*asynqserver.Response, error) {
	// todo: add your logic here and delete this line

	return &asynqserver.Response{}, nil
}
