package sentry

import (
	"os"
	"strconv"

	"github.com/getsentry/raven-go"
)

// 捕获事件级别的
func Capture(packet *raven.Packet, captureTags map[string]string) (eventID string, ch chan error) {
	return raven.Capture(packet, captureTags)
}

func CaptureError(err error, tags map[string]string, interfaces ...raven.Interface) string {
	return raven.CaptureError(err, tags, interfaces...)
}

// 带Ack的捕获错误
func CaptureErrorAndWait(err error, tags map[string]string, interfaces ...raven.Interface) string {
	return raven.CaptureErrorAndWait(err, tags, interfaces...)
}

func CapturePanic(f func(), tags map[string]string, interfaces ...raven.Interface) (err interface{}, errorID string) {
	return raven.CapturePanic(f, tags, interfaces...)
}

// 带Ack的捕获panic
func CapturePanicAndWait(f func(), tags map[string]string, interfaces ...raven.Interface) (err interface{}, errorID string) {
	return raven.CapturePanicAndWait(f, tags, interfaces...)
}

func CaptureMessage(message string, tags map[string]string, interfaces ...raven.Interface) string {
	return raven.CaptureMessage(message, tags, interfaces...)
}

// 带Ack的捕获message
func CaptureMessageAndWait(message string, tags map[string]string, interfaces ...raven.Interface) string {
	return raven.CaptureMessageAndWait(message, tags, interfaces...)
}

// sentry tags
func AllPropertyforSentry(meta *RequestMeta, path string) map[string]string {
	return map[string]string{
		"client":        meta.AppName,
		"platform":      meta.Platform,
		"language":      meta.LanguageApp,
		"deviceType":    meta.DeviceType,
		"deviceDensity": meta.Resolution,
		"countryCode":   meta.Country,
		"appVersion":    meta.AppVersion,
		"network":       meta.NetType,
		"userId":        strconv.FormatInt(meta.UserId, 10),
		"userIP":        meta.UserIp,
		"realUserIP":    meta.UserIp,
		"X-Request-id":  strconv.FormatInt(meta.TraceId, 10),
		"URL_PATH":      path,
		"server_name":   GetHostName(),
	}
}

func GetHostName() string {
	if h, err := os.Hostname(); err == nil {
		return h
	}
	return ""
}
type RequestMeta struct {
	TraceId int64 `protobuf:"varint,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Android/ios
	Platform string `protobuf:"bytes,3,opt,name=platform" json:"platform,omitempty"`
	AppName  string `protobuf:"bytes,4,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	// 7.5.6
	AppVersion       string   `protobuf:"bytes,5,opt,name=app_version,json=appVersion" json:"app_version,omitempty"`
	LanguageApp      string   `protobuf:"bytes,6,opt,name=language_app,json=languageApp" json:"language_app,omitempty"`
	LanguageContent  string   `protobuf:"bytes,7,opt,name=language_content,json=languageContent" json:"language_content,omitempty"`
	DeviceId         string   `protobuf:"bytes,8,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	UserIp           string   `protobuf:"bytes,9,opt,name=user_ip,json=userIp" json:"user_ip,omitempty"`
	SkipCheck        int32    `protobuf:"varint,10,opt,name=skip_check,json=skipCheck" json:"skip_check,omitempty"`
	Country          string   `protobuf:"bytes,11,opt,name=country" json:"country,omitempty"`
	Province         string   `protobuf:"bytes,12,opt,name=province" json:"province,omitempty"`
	ContentLanguages []string `protobuf:"bytes,13,rep,name=content_languages,json=contentLanguages" json:"content_languages,omitempty"`
	// wifi
	NetType    string `protobuf:"bytes,14,opt,name=net_type,json=netType" json:"net_type,omitempty"`
	Resolution string `protobuf:"bytes,15,opt,name=resolution" json:"resolution,omitempty"`
	DeviceType string `protobuf:"bytes,16,opt,name=device_type,json=deviceType" json:"device_type,omitempty"`
	CountrySim string `protobuf:"bytes,17,opt,name=country_sim,json=countrySim" json:"country_sim,omitempty"`
	// android 10.0
	OsVersion string `protobuf:"bytes,18,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	// api host
	Host string `protobuf:"bytes,19,opt,name=host" json:"host,omitempty"`
	// 品牌
	PhoneBrand string `protobuf:"bytes,20,opt,name=phone_brand,json=phoneBrand" json:"phone_brand,omitempty"`
	// 型号
	PhoneModel string `protobuf:"bytes,21,opt,name=phone_model,json=phoneModel" json:"phone_model,omitempty"`
	// 厂商
	PhoneManufacturer string `protobuf:"bytes,22,opt,name=phone_manufacturer,json=phoneManufacturer" json:"phone_manufacturer,omitempty"`
}