package stringsvc_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func TestHandlers(t *testing.T) {
	cases := []struct {
		method, route    string
		reqBody          string
		expectedRespBody string
		expectedRespCode int
	}{
		{
			method:           "POST",
			route:            "/uppercase",
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":"HELLO, WORLD"}`,
			expectedRespCode: http.StatusOK,
		},
		{
			method:           "POST",
			route:            "/uppercase",
			reqBody:          `{"s":""}`,
			expectedRespBody: `{"v":"","err":"Empty string"}`,
			expectedRespCode: http.StatusOK,
		},
		{
			method:           "POST",
			route:            "/count",
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":12}`,
			expectedRespCode: http.StatusOK,
		},
	}

	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	s := stringsvc.NewStringService()
	h := stringsvc.MakeHTTPHandler(ctx, s, logger)

	for _, c := range cases {
		t.Run(
			fmt.Sprintf("route=%q,method=%q,body=%s", c.route, c.method, c.reqBody),
			func(t *testing.T) {
				req, _ := http.NewRequest(c.method, c.route, strings.NewReader(c.reqBody))
				rec := httptest.NewRecorder()
				h.ServeHTTP(rec, req)

				errMsg := "%s %s, body: %s - want %v, got %v"

				if rec.Code != c.expectedRespCode {
					t.Errorf(errMsg, c.method, c.route, c.reqBody, c.expectedRespCode, rec.Code)
				}

				respBody := strings.TrimSpace(rec.Body.String())
				if respBody != c.expectedRespBody {
					t.Errorf(errMsg, c.method, c.route, c.reqBody, c.expectedRespBody, respBody)
				}
			})
	}
}

func BenchmarkHandlers(b *testing.B) {
	cases := []struct {
		method, route string
		reqBody       string
	}{
		{
			method:  "POST",
			route:   "/uppercase",
			reqBody: `{"s":"hello, world"}`,
		},
		{
			method:  "POST",
			route:   "/uppercase",
			reqBody: `{"s":""}`,
		},
		{
			method:  "POST",
			route:   "/count",
			reqBody: `{"s":"hello, world"}`,
		},
	}

	ctx := context.Background()

	// Discard log output
	logger := log.NewLogfmtLogger(log.NewSyncWriter(ioutil.Discard))

	var s stringsvc.Service
	{
		s = stringsvc.NewStringService()
		s = stringsvc.LoggingMiddleware(logger)(s)
	}

	h := stringsvc.MakeHTTPHandler(ctx, s, logger)
	rec := httptest.NewRecorder()

	for _, c := range cases {
		b.Run(
			fmt.Sprintf("route=%q,method=%q,body=%s", c.route, c.method, c.reqBody),
			func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					req, _ := http.NewRequest(c.method, c.route, strings.NewReader(c.reqBody))
					h.ServeHTTP(rec, req)
				}
			})
	}
}
