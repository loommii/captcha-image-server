package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Captcha CaptchaConfig
}
type CaptchaConfig struct {
	Expire int
}
