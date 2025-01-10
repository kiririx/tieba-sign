# tieba-sign
贴吧手机端签到

### 相关接口参考 https://github.com/LuoSue/TiebaSignIn-1

## 概述
利用贴吧手机端api，通过BDUSS （Cookie中的一段字符串）来实现自动签到, 支持结果推送到pushdeer

## 使用
### 命令行执行：
#### Windows
```
set bduss=你的bduss
# 如果有pushdeer地址就填写, 没有就不写
set pushdeer.addr=https://abc.com/message/push
set pushdeer.pushkey=你的pushkey
go run main.go
```

#### Linux & MacOS
```
export bduss=你的bduss
# 如果有pushdeer地址就填写, 没有就不写
export pushdeer.addr=https://abc.com/message/push
export pushdeer.pushkey=你的pushkey
go run main.go
```

### docker
```
docker run --rm -e bduss=你的bduss -e pushdeer.pushkey=你的pushkey -e pushdeer.addr=你的pushdeer地址 kiririx/tieba_sign:latest
```

## 获取BDUSS
通过在电脑端浏览器登入贴吧，打开开发者工具，查看Cookie中的BDUSS。
