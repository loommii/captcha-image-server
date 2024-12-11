package logic

import (
	"context"

	"captchaImageServer/internal/svc"
	"captchaImageServer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaImageServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaImageServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaImageServerLogic {
	return &CaptchaImageServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaImageServerLogic) CaptchaImageServer(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
