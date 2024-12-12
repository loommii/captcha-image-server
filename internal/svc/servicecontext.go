package svc

import (
	"captchaImageServer/internal/config"
	"time"

	"github.com/loommii/captcha-image/captchaimage"
	"github.com/zeromicro/go-zero/core/collection"
)

type ServiceContext struct {
	Config       config.Config
	Captcha      *captchaimage.CaptchaGenerator
	CaptchaCache *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	captchaCache, err := collection.NewCache(time.Second * time.Duration(c.Captcha.Expire))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:       c,
		Captcha:      captchaimage.NewCaptchaGenerator(),
		CaptchaCache: captchaCache,
	}
}
