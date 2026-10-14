package pkg

import (
	"testing"
)

func TestLoginURL(t *testing.T) {
	t.Setenv("FI_MCP_PUBLIC_URL", "")
	t.Setenv("FI_MCP_PORT", "9090")

	got := LoginURL("session/id with spaces")
	want := "http://localhost:9090/mockWebPage?sessionId=session%2Fid+with+spaces"
	if got != want {
		t.Fatalf("LoginURL() = %q, want %q", got, want)
	}

	t.Setenv("FI_MCP_PUBLIC_URL", "https://fi.example/mcp/")
	got = LoginURL("abc")
	want = "https://fi.example/mcp/mockWebPage?sessionId=abc"
	if got != want {
		t.Fatalf("LoginURL() with public URL = %q, want %q", got, want)
	}
}

func TestResolvePhone(t *testing.T) {
	allowed := []string{"1111111111", "2222222222"}

	t.Run("session phone wins", func(t *testing.T) {
		t.Setenv("FI_MCP_LOGIN_PHONE", "2222222222")
		phone, ok := ResolvePhone("1111111111", true, allowed)
		if !ok || phone != "1111111111" {
			t.Fatalf("got %q %v", phone, ok)
		}
	})

	t.Run("unset pin asks for login", func(t *testing.T) {
		t.Setenv("FI_MCP_LOGIN_PHONE", "")
		if _, ok := ResolvePhone("", false, allowed); ok {
			t.Fatal("expected login to be required")
		}
	})

	t.Run("allowed pin", func(t *testing.T) {
		t.Setenv("FI_MCP_LOGIN_PHONE", "2222222222")
		phone, ok := ResolvePhone("", false, allowed)
		if !ok || phone != "2222222222" {
			t.Fatalf("got %q %v", phone, ok)
		}
	})

	t.Run("unknown pin is ignored", func(t *testing.T) {
		t.Setenv("FI_MCP_LOGIN_PHONE", "0000000000")
		if _, ok := ResolvePhone("", false, allowed); ok {
			t.Fatal("expected unknown pin to be ignored")
		}
	})
}
