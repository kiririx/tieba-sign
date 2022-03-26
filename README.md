# tieba-sign
贴吧手机端签到

### 相关接口参考 https://github.com/LuoSue/TiebaSignIn-1

## 概述
利用贴吧手机端api，通过BDUSS （Cookie中的一段字符串）来实现自动签到，目前尚在开发中。

## 使用
直接编译部署，通过访问前端页面，添加自己的BDUSS，也可以设置每天定时签到；

通过前端页面可以查看签到状态，签到次数等。

## 技术
- Core：Go、Gin、Gorm
- UI：React
- Config：MySQL、配置文件

## 配置项
- mysql.conn.user=用户名
- mysql.conn.pass=密码
- mysql.conn.database=数据库
- mysql.conn.ip=xxx.xxx.xxx.xxx
- mysql.conn.port=3306
- task.start.hour=6 （每天开启自动签到的时间点，0-23）

## 获取BDUSS
通过在电脑端浏览器登入贴吧，打开开发者工具，查看Cookie中的BDUSS。