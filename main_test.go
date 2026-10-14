package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/epifi/fi-mcp-lite/middlewares"
)

func postLogin(t *testing.T, phone string) *httptest.ResponseRecorder {
	t.Helper()
	form := url.Values{}
	form.Set("sessionId", "session-test")
	form.Set("phoneNumber", phone)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	loginHandler(rec, req)
	return rec
}

func TestLoginHandlerRejectsUnknownPhone(t *testing.T) {
	authMiddleware = nil
	rec := postLogin(t, "0000000000")
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "phone number is not allowed") {
		t.Fatalf("body = %q", rec.Body.String())
	}
}

func TestLoginHandlerAcceptsPersonaPhone(t *testing.T) {
	authMiddleware = middlewares.NewAuthMiddleware()
	rec := postLogin(t, "1111111111")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %q", rec.Code, rec.Body.String())
	}
}
