package sentry

import "github.com/getsentry/raven-go"

const (
	prodDSN = "http://aaa:f43bcdd21de14cc5ac7c946bc78eecb6@sentry.aaa.com/2"
	boxDSN  = "http://bbbb:400a78a15868434c8b1c219e120890f0@sentry.aaa.com/7"
)

func InitSentry(isTest bool, env string) {
	InitDsn(isTest)
	InitEnv(env)
}

func InitDsn(isTest bool) {
	dsn := prodDSN
	if isTest {
		dsn = boxDSN
	}
	raven.SetDSN(dsn)
}

func InitEnv(env string) {
	raven.SetEnvironment(env)
}

// 采样比例
func SetRate(rate float32) {
	raven.SetSampleRate(rate)
}

func InitOther() {
	raven.SetDefaultLoggerName("some_logger_name")
}

// 在发送到 Sentry 之前要过滤掉的消息列表。这个列表将形成一个 RegExp，它将检查错误消息或用户直接传递的消息的部分匹配。
func InitIgnore() {
	raven.SetIgnoreErrors([]string{"ThirdPartyServiceUnavailable", "Other error that we want to ignore"}...)
}

// 属于应用程序的模块名称的字符串前缀列表。此选项将用于确定 Sentry SDK 是否应将框架标记为用户或本机/外部代码。
func InitIncludePath() {
	raven.SetIncludePaths([]string{"/some/path", "other/path"})
}
