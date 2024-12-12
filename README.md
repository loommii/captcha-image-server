 <div align="center">
  <a href="https://github.com/loommii/captcha-image-server"><img width="150 px" alt="logo" src="captcha-image-server.png"/></a>
  <p><em>验证码图片生成</em></p>
</div>

-----------

# `captcha-image-server`
>这是一个基于Go-Zero开发的验证码校验服务
- 图片验证码生成
- 验证码校验


### Docker 部署
```
docker run -d -p 54289:54289 --name captcha-image-server loommii/captcha-image-server
```

### DEMO服务-CURL
```
# Ping接口
curl --location --request GET 'https://captcha.1236677.xyz/ping'

# 获取验证码图片 - 自行生成UUID替换
curl --location --request GET 'https://captcha.1236677.xyz/captcha/image?id=uuid'

# 校验验证码 - 获取验证码图片返回的ID
curl --location --request POST 'https://captcha.1236677.xyz/captcha/Check' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "uuid",
    "value": "验证码"
}'
```