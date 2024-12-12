package logic

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"image/png"
	"time"

	"captchaImageServer/internal/svc"
	"captchaImageServer/internal/types"

	"github.com/duke-git/lancet/v2/random"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaImageLogic {
	return &GetCaptchaImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaImageLogic) GetCaptchaImage(req *types.GetCaptchaImageReq) (resp *types.GetCaptchaImageResp, err error) {
	// 重复ID错误返回
	_, ok := l.svcCtx.CaptchaCache.Get(req.ID)
	if ok {
		return nil, errors.New("重复ID")
	}
	captcha := random.RandString(6)                       // 随机生成验证码
	img, err := l.svcCtx.Captcha.GenerateCaptcha(captcha) // 生成图片
	if err != nil {
		return
	}
	// 将验证码缓存
	l.svcCtx.CaptchaCache.Set(req.ID, captcha)
	// 图片返回
	var buf bytes.Buffer
	png.Encode(&buf, img)

	resp = &types.GetCaptchaImageResp{}
	resp.ID = req.ID
	resp.ExpireUnix = int(time.Now().Unix()) + l.svcCtx.Config.Captcha.Expire
	resp.CaptchaImageB64 = base64.RawStdEncoding.EncodeToString(buf.Bytes())
	return
}
