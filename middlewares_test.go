package stringsvc_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

type service struct {
	UppercaseF func(context.Context, string) (string, error)
	CountF     func(context.Context, string) int
}

func (svc *service) Uppercase(ctx context.Context, s string) (string, error) {
	return svc.UppercaseF(ctx, s)
}

func (svc *service) Count(ctx context.Context, s string) (n int) {
	return svc.CountF(ctx, s)
}

var (
	logfmtRegex = `^method=(\S+) ` +
		`input=("*[a-zA-Z0-9_ ]*"*) ` +
		`output=("*[a-zA-Z0-9_ ]*"*) ` +
		`err=("*[a-zA-Z0-9_ ]*"*) ` +
		`took=(\S+)$`
	methodPrefix = `method=`
)

func TestLogMiddlewareUppercase(t *testing.T) {
	ctx := context.Background()

	var logBuf bytes.Buffer
	wr := io.MultiWriter(&logBuf, os.Stderr)
	logger := log.NewLogfmtLogger(wr)

	svc := &service{}
	logSvc := stringsvc.LoggingMiddleware(logger)(svc)

	svc.UppercaseF = func(ctx context.Context, s string) (string, error) {
		return "", nil
	}
	logSvc.Uppercase(ctx, "")

	testLogFmt(t, logBuf, "uppercase")
}

func TestLogMiddlewareCount(t *testing.T) {
	ctx := context.Background()

	var logBuf bytes.Buffer
	wr := io.MultiWriter(&logBuf, os.Stderr)
	logger := log.NewLogfmtLogger(wr)

	svc := &service{}
	logSvc := stringsvc.LoggingMiddleware(logger)(svc)

	svc.CountF = func(ctx context.Context, s string) int {
		return 0
	}
	logSvc.Count(ctx, "")

	testLogFmt(t, logBuf, "count")
}

func testLogFmt(t *testing.T, logOutput bytes.Buffer, method string) {
	b, err := ioutil.ReadAll(&logOutput)
	if err != nil {
		t.Error(err)
	}
	b = bytes.TrimSpace(b)
	match, err := regexp.Match(logfmtRegex, b)
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Errorf("log output does not match regex %q, ouput %q", logfmtRegex, b)
	}
	prefix := []byte(methodPrefix + method)
	if !bytes.HasPrefix(b, prefix) {
		t.Errorf("log output method does not match, want prefix %q, got %q", prefix, b)
	}
}
