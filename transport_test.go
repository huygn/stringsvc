package stringsvc

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/context"
)

func TestStringUppercase(t *testing.T) {
	const (
		url    = "/uppercase"
		method = "POST"
		code   = http.StatusOK
	)
	reqBody := `{"s":"hello, world"}`
	expectedResp := `{"v":"HELLO, WORLD"}`

	ctx := context.Background()
	s := NewStringService()
	h := MakeHTTPHandler(ctx, s)

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

func TestStringCount(t *testing.T) {
	const (
		url    = "/count"
		method = "POST"
		code   = http.StatusOK
	)
	reqBody := `{"s":"hello, world"}`
	expectedResp := `{"v":12}`

	ctx := context.Background()
	s := NewStringService()
	h := MakeHTTPHandler(ctx, s)

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
