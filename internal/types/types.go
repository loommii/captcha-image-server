// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type CaptchaCheckReq struct {
	ID    string `json:"id"`    // 验证码对应的ID
	Value string `json:"value"` // 验证码
}

type CaptchaCheckResp struct {
	ID       string `json:"id"`       // 验证码对应的ID
	Result   bool   `json:"result"`   // 成功或失败
	Redirect string `json:"redirect"` // 成功的情况下,应该要让用户访问。类重定向
}

type GetCaptchaImageReq struct {
	ID string `form:"id"` // 建议请求方用UUID
}

type GetCaptchaImageResp struct {
	ID              string `json:"id"`              // 请求的UUID
	ExpireUnix      int    `json:"expireUnix"`      // 过期时间戳
	CaptchaImageB64 string `json:"captchaImageB64"` // 验证码图片B64
}

type PingResp struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}
