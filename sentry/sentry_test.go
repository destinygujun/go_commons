package sentry

import (
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"testing"
)

func TestCapturePanic(t *testing.T) {
	meta := new(RequestMeta) // GetRequestMeta(c)
	path := ""                     // c.Request.URL.String()
	CapturePanic(func() {
		panic(fmt.Sprintf("[Recovery] panic recovered: %s", "error"))
	}, AllPropertyforSentry(meta, path))
}

func TestCaptureError(t *testing.T) {
	meta := new(RequestMeta) // GetRequestMeta(c)
	path := ""                     // c.Request.URL.String()
	err := errors.New("item not found")
	CaptureError(err, AllPropertyforSentry(meta, path))
	CaptureErrorAndWait(err, AllPropertyforSentry(meta, path))
}

func TestCaptureMessage(t *testing.T) {
	meta := new(RequestMeta) // GetRequestMeta(c)
	path := ""                     // c.Request.URL.String()
	msg := "item not found"
	CaptureMessage(msg, AllPropertyforSentry(meta, path))
	CaptureMessageAndWait(msg, AllPropertyforSentry(meta, path))
}

func TestCapture(t *testing.T) {
	meta := new(RequestMeta) // GetRequestMeta(c)
	path := ""                     // c.Request.URL.String()
	msg := "item not found"
	pack := &raven.Packet{
		Message: msg,
	}
	Capture(pack, AllPropertyforSentry(meta, path))
}
