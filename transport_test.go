package stringsvc_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/gnhuy91/stringsvc"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		method, url               string
		code                      int
		reqBody, expectedRespBody string
	}{
		{
			method:           "POST",
			url:              "/uppercase",
			code:             http.StatusOK,
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":"HELLO, WORLD"}`,
		},
		{
			method:           "POST",
			url:              "/count",
			code:             http.StatusOK,
			reqBody:          `{"s":"hello, world"}`,
			expectedRespBody: `{"v":12}`,
		},
	}

	for _, c := range cases {
		ctx := context.Background()
		logger := log.NewLogfmtLogger(os.Stderr)
		s := stringsvc.NewStringService()
		h := stringsvc.MakeHTTPHandler(ctx, s, logger)

		req, _ := http.NewRequest(c.method, c.url, strings.NewReader(c.reqBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		errMsg := "%s %s, body: %s - want %v, got %v"

		if rec.Code != c.code {
			t.Errorf(errMsg, c.method, c.url, c.reqBody, c.code, rec.Code)
		}

		respBody := strings.TrimSpace(rec.Body.String())
		if respBody != c.expectedRespBody {
			t.Errorf(errMsg, c.method, c.url, c.reqBody, c.expectedRespBody, respBody)
		}
	}
}
