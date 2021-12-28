#sentry引入

##1. sentry基本信息

###1.1 可视化
http://sentry.ushow.media/qmkl/sm-api

###1.2 doc
https://docs.sentry.io/clients/go/

##2. 使用步骤
###2.1 初始化sdk
go get github.com/getsentry/raven-go

###2.2 项目配置
raven.SetDSN(dsn)
如：./config.go中的 prodDSN（待提取到配置文件）

###2.3 项目引入
启动文件初始化  
如：api_gateway/main.go
sentry.InitSentry(config.IsTest(), config.GetEnv())

###2.4 实际使用
###2.4.1 捕获panic
如：api_gateway/proxy.go:125 recover上报panic
如./sentry_test.go TestCapturePanic
tag可聚合自定义

###2.4.2 捕获error
如./sentry_test.go TestCaptureError
tag可聚合自定义

###2.4.3 捕获message
如./sentry_test.go TestCaptureMessage
tag可聚合自定义

###2.4.4 捕获事件
如./sentry_test.go TestCapture
事件具体字段可自定义，tag可聚合自定义

##2.5 其他模块待引入
……