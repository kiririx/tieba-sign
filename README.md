# tieba-sign
贴吧手机端签到

### 相关接口参考 https://github.com/LuoSue/TiebaSignIn-1

## 概述
利用贴吧手机端api，通过BDUSS （Cookie中的一段字符串）来实现自动签到

## 使用
命令行执行：./run -h 12 -b your_bduss

run 表示你的二进制程序文件

-h 表示自动签到的小时（0-23）

-b 表示你的bduss值

## 获取BDUSS
通过在电脑端浏览器登入贴吧，打开开发者工具，查看Cookie中的BDUSS。