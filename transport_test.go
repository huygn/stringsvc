package stringsvc_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	. "github.com/gnhuy91/stringsvc"
)

func TestHandlerUppercase(t *testing.T) {
	const (
		url    = "/uppercase"
		method = "POST"
		code   = http.StatusOK
	)
	reqBody := `{"s":"hello, world"}`
	expectedResp := `{"v":"HELLO, WORLD"}`

	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	s := NewStringService()
	h := MakeHTTPHandler(ctx, s, logger)

	req, _ := http.NewRequest(method, url, strings.NewReader(reqBody))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)
	errMsg := "%s %s, body: %s - want %v, got %v"

	if rec.Code != code {
		t.Errorf(errMsg, method, url, reqBody, code, rec.Code)
	}

	respBody := strings.TrimSpace(rec.Body.String())
	if respBody != expectedResp {
		t.Errorf(errMsg, method, url, reqBody, expectedResp, respBody)
	}
}

func TestHandlerCount(t *testing.T) {
	const (
		url    = "/count"
		method = "POST"
		code   = http.StatusOK
	)
	reqBody := `{"s":"hello, world"}`
	expectedResp := `{"v":12}`

	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	s := NewStringService()
	h := MakeHTTPHandler(ctx, s, logger)

	req, _ := http.NewRequest(method, url, strings.NewReader(reqBody))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)
	errMsg := "%s %s, body: %s - want %v, got %v"

	if rec.Code != code {
		t.Errorf(errMsg, method, url, reqBody, code, rec.Code)
	}

	respBody := strings.TrimSpace(rec.Body.String())
	if respBody != expectedResp {
		t.Errorf(errMsg, method, url, reqBody, expectedResp, respBody)
	}
}
