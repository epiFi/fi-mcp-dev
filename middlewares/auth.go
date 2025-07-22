package middlewares

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/samber/lo"

	"github.com/epifi/fi-mcp-lite/pkg"
)

var (
	loginRequiredJson = `{"status": "login_required","login_url": "%s","message": "Needs to login first by going to the login url.\nShow the login url as clickable link if client supports it. Otherwise display the URL for users to copy and paste into a browser. \nAsk users to come back and let you know once they are done with login in their browser"}`
)

type AuthedUser struct {
	PhoneNumber string
	LoginTime   time.Time
}

type AuthMiddleware struct {
	sessionStore map[string]string // session -> phone number (existing)
	userAuth     map[string]*AuthedUser // phone number -> auth info (new)
	authDuration time.Duration
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		sessionStore: make(map[string]string),
		userAuth:     make(map[string]*AuthedUser),
		authDuration: 30 * time.Minute, // Auth persists for 30 minutes
	}
}

func (m *AuthMiddleware) AuthMiddleware(next server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// fetch sessionId from context
		// this gets populated for every tool call
		sessionId := server.ClientSessionFromContext(ctx).SessionID()
		
		// Debug logging
		log.Printf("DEBUG: Tool call sessionId: %s", sessionId)
		log.Printf("DEBUG: Stored sessions: %v", m.sessionStore)
		log.Printf("DEBUG: Stored user auth: %v", m.userAuth)
		
		// First, check if this specific session is authenticated
		phoneNumber, ok := m.sessionStore[sessionId]
		if !ok {
			// If session not found, check if any user has recent valid auth
			phoneNumber = m.findRecentlyAuthedUser()
			if phoneNumber == "" {
				log.Printf("DEBUG: No valid auth found for sessionId: %s", sessionId)
				loginUrl := m.getLoginUrl(sessionId)
				return mcp.NewToolResultText(fmt.Sprintf(loginRequiredJson, loginUrl)), nil
			}
			// Store this session for the authenticated user
			m.sessionStore[sessionId] = phoneNumber
			log.Printf("DEBUG: Using recent auth for user %s, stored session %s", phoneNumber, sessionId)
		}
		
		log.Printf("DEBUG: Found session for sessionId: %s, phoneNumber: %s", sessionId, phoneNumber)
		
		if !lo.Contains(pkg.GetAllowedMobileNumbers(), phoneNumber) {
			return mcp.NewToolResultError("phone number is not allowed"), nil
		}
		ctx = context.WithValue(ctx, "phone_number", phoneNumber)
		toolName := req.Params.Name
		data, readErr := os.ReadFile("test_data_dir/" + phoneNumber + "/" + toolName + ".json")
		if readErr != nil {
			log.Println("error reading test data file", readErr)
			return mcp.NewToolResultError("error reading test data file"), nil
		}
		return mcp.NewToolResultText(string(data)), nil
	}
}

// findRecentlyAuthedUser checks if any user has authenticated recently
func (m *AuthMiddleware) findRecentlyAuthedUser() string {
	now := time.Now()
	for phoneNumber, authInfo := range m.userAuth {
		if now.Sub(authInfo.LoginTime) < m.authDuration {
			log.Printf("DEBUG: Found recent auth for user %s, logged in %v ago", phoneNumber, now.Sub(authInfo.LoginTime))
			return phoneNumber
		}
	}
	return ""
}

// GetLoginUrl fetches dynamic login url for given sessionId
func (m *AuthMiddleware) getLoginUrl(sessionId string) string {
	return fmt.Sprintf("http://localhost:%s/mockWebPage?sessionId=%s", pkg.GetPort(), sessionId)
}

func (m *AuthMiddleware) AddSession(sessionId, phoneNumber string) {
	log.Printf("DEBUG: AddSession called - sessionId: %s, phoneNumber: %s", sessionId, phoneNumber)
	
	// Store session mapping
	m.sessionStore[sessionId] = phoneNumber
	
	// Store user auth with timestamp
	m.userAuth[phoneNumber] = &AuthedUser{
		PhoneNumber: phoneNumber,
		LoginTime:   time.Now(),
	}
	
	log.Printf("DEBUG: Session stored. Current store: %v", m.sessionStore)
	log.Printf("DEBUG: User auth stored. Current auth: %v", m.userAuth)
}
