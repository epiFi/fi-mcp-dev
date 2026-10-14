package pkg

import (
	"log"
	"net/url"
	"os"
	"strings"
)

// PublicBaseURL is the origin used in login links returned to MCP clients.
// FI_MCP_PUBLIC_URL overrides the default so Docker, Codespaces, and tunnels
// get a link the user can actually open. A localhost link on a remote host
// is one reason clients appear to ask for login forever.
func PublicBaseURL() string {
	if base := strings.TrimSpace(os.Getenv("FI_MCP_PUBLIC_URL")); base != "" {
		return strings.TrimRight(base, "/")
	}
	return "http://localhost:" + GetPort()
}

// LoginURL is the mock login page for an MCP session.
func LoginURL(sessionID string) string {
	return PublicBaseURL() + "/mockWebPage?sessionId=" + url.QueryEscape(sessionID)
}

// DevLoginPhone is the optional FI_MCP_LOGIN_PHONE persona. Empty means the
// normal browser login is required.
func DevLoginPhone() string {
	return strings.TrimSpace(os.Getenv("FI_MCP_LOGIN_PHONE"))
}

// ResolvePhone returns the phone already bound to a session, or the pinned
// dev persona when the session has none. The pinned phone is used only when
// it is one of the dummy personas passed in allowed.
func ResolvePhone(sessionPhone string, found bool, allowed []string) (string, bool) {
	if found && sessionPhone != "" {
		return sessionPhone, true
	}

	phone := DevLoginPhone()
	if phone == "" {
		return "", false
	}

	for _, candidate := range allowed {
		if candidate == phone {
			return phone, true
		}
	}

	log.Printf("FI_MCP_LOGIN_PHONE %q is not an allowed test persona", phone)
	return "", false
}
