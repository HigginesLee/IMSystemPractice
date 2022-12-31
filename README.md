# 需求分析

- 消息发送与接收
- 访客、点对点、群聊、广播、快捷恢复、撤回、心跳检测

## 技术栈

- websocket 组件转发消息
- channel/goroutine 提高并发性
- gin/template/swagger
- gorm, logger, govalidator
- SQL/NOSQL/MQ 等

## 系统架构:

四层:前端、接入层、业务逻辑处理、存储

消息发送流程:
A> 登录> 鉴权(游客)> 消息类型(单/群/广播)>B
