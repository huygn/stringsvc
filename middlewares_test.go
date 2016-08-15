package stringsvc_test

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "")
	flag.Parse()
}

var logfmtRegex = `^method=(\S+) ` +
	`input=("*.*"*) ` +
	`output=("*.*"*) ` +
	`err=("*.*"*) ` +
	`took=(\S+)$`

func TestLogMiddlewareUppercase(t *testing.T) {
	tc := struct {
		input, output string
	}{
		input:  `hello, world`,
		output: `method=uppercase input="hello, world" output="HELLO, WORLD" err=null`,
	}

	ctx := context.Background()

	var buf bytes.Buffer
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(&buf)
		if debug {
			logger = log.NewLogfmtLogger(io.MultiWriter(&buf, os.Stderr))
		}
	}

	svc := stringsvc.NewStringService()
	logMw := stringsvc.LoggingMiddleware(logger)(svc)
	logMw.Uppercase(ctx, tc.input)

	testLogFmt(t, buf, tc.output)
}

func TestLogMiddlewareCount(t *testing.T) {
	tc := struct {
		input, output string
	}{
		input:  `hello, world`,
		output: `method=count input="hello, world" output=12 err=null`,
	}

	ctx := context.Background()

	var buf bytes.Buffer
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(&buf)
		if debug {
			logger = log.NewLogfmtLogger(io.MultiWriter(&buf, os.Stderr))
		}
	}

	svc := stringsvc.NewStringService()
	logMw := stringsvc.LoggingMiddleware(logger)(svc)
	logMw.Count(ctx, tc.input)

	testLogFmt(t, buf, tc.output)
}

func testLogFmt(t *testing.T, logOutput bytes.Buffer, prefix string) {
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
		t.Errorf("log output does not match regex %s, ouput %s", logfmtRegex, b)
	}
	p := []byte(prefix)
	if !bytes.HasPrefix(b, p) {
		t.Errorf("log output method does not match, want prefix %q, got %q", p, b)
	}
}
