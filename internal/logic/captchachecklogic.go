package logic

import (
	"context"

	"captchaImageServer/internal/svc"
	"captchaImageServer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaCheckLogic {
	return &CaptchaCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaCheckLogic) CaptchaCheck(req *types.CaptchaCheckReq) (resp *types.CaptchaCheckResp, err error) {
	resp = &types.CaptchaCheckResp{}
	resp.ID = req.ID
	// 重复ID错误返回
	v, ok := l.svcCtx.CaptchaCache.Get(req.ID)
	if !ok || v != req.Value { // ID不存在或验证码不正确的情况
		return
	}
	resp.Result = true
	resp.Redirect = "https://loommii.github.io/"
	// TODO: 其他业务逻辑
	return
}
