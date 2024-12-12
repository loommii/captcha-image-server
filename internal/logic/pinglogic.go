package logic

import (
	"context"
	"time"

	"captchaImageServer/internal/svc"
	"captchaImageServer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.PingResp, err error) {
	resp = &types.PingResp{
		Message: "ok",
		Time:    time.Now().Local().String(),
	}
	return
}
