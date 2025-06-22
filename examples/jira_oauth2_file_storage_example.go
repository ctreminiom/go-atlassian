package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// FileTokenStore implements oauth2.TokenStore using filesystem
type FileTokenStore struct {
	basePath string
	mu       sync.RWMutex
}

// NewFileTokenStore creates a new file-based token store
func NewFileTokenStore(basePath string) (*FileTokenStore, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(basePath, 0700); err != nil {
		return nil, fmt.Errorf("failed to create token store directory: %w", err)
	}
	
	return &FileTokenStore{
		basePath: basePath,
	}, nil
}

// GetToken retrieves the OAuth2 token from file
func (f *FileTokenStore) GetToken(ctx context.Context) (*common.OAuth2Token, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	
	tokenPath := filepath.Join(f.basePath, "token.json")
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("token not found")
		}
		return nil, fmt.Errorf("failed to read token: %w", err)
	}
	
	var token common.OAuth2Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}
	
	return &token, nil
}

// SetToken stores the OAuth2 token to file
func (f *FileTokenStore) SetToken(ctx context.Context, token *common.OAuth2Token) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	
	data, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}
	
	tokenPath := filepath.Join(f.basePath, "token.json")
	if err := os.WriteFile(tokenPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write token: %w", err)
	}
	
	return nil
}

// GetRefreshToken retrieves only the refresh token from file
func (f *FileTokenStore) GetRefreshToken(ctx context.Context) (string, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	
	refreshPath := filepath.Join(f.basePath, "refresh_token")
	data, err := os.ReadFile(refreshPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("refresh token not found")
		}
		return "", fmt.Errorf("failed to read refresh token: %w", err)
	}
	
	return string(data), nil
}

// SetRefreshToken stores only the refresh token to file
func (f *FileTokenStore) SetRefreshToken(ctx context.Context, refreshToken string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	
	refreshPath := filepath.Join(f.basePath, "refresh_token")
	if err := os.WriteFile(refreshPath, []byte(refreshToken), 0600); err != nil {
		return fmt.Errorf("failed to write refresh token: %w", err)
	}
	
	return nil
}

// AuditTokenCallback implements oauth2.TokenCallback for audit logging
type AuditTokenCallback struct {
	logFile string
	mu      sync.Mutex
}

func NewAuditTokenCallback(logFile string) *AuditTokenCallback {
	return &AuditTokenCallback{
		logFile: logFile,
	}
}

func (a *AuditTokenCallback) OnTokenRefreshed(ctx context.Context, oldToken, newToken *common.OAuth2Token) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	
	// Open or create audit log file
	file, err := os.OpenFile(a.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Write audit entry
	entry := fmt.Sprintf("[%s] Token refreshed. New token expires in %d seconds\n", 
		time.Now().Format(time.RFC3339), newToken.ExpiresIn)
	_, err = file.WriteString(entry)
	
	return err
}

// Example: Using OAuth 2.0 with file-based token storage
func ExampleOAuth2WithFileStorage() {
	// Create file-based token store
	homeDir, _ := os.UserHomeDir()
	tokenStorePath := filepath.Join(homeDir, ".atlassian-oauth")
	
	tokenStore, err := NewFileTokenStore(tokenStorePath)
	if err != nil {
		log.Fatalf("Failed to create token store: %v", err)
	}
	
	fmt.Printf("Token store created at: %s\n", tokenStorePath)
	
	// OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Try to load existing token from file
	var existingToken *common.OAuth2Token
	storedToken, err := tokenStore.GetToken(context.Background())
	if err == nil {
		existingToken = storedToken
		fmt.Println("Loaded existing token from file")
	} else {
		// If no token in file, use a new one (in real app, this would come from OAuth flow)
		existingToken = &common.OAuth2Token{
			AccessToken:  "existing-access-token",
			TokenType:    "Bearer",
			ExpiresIn:    3600, // 1 hour
			RefreshToken: "existing-refresh-token",
			Scope:        "read:jira-work write:jira-work offline_access",
		}
		fmt.Println("Using new token (not found in file)")
	}
	
	// Create audit callback
	auditLog := filepath.Join(tokenStorePath, "token_audit.log")
	auditCallback := NewAuditTokenCallback(auditLog)
	
	// Create client with file storage and audit logging
	client, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithTokenStore(tokenStore),
		jira.WithTokenCallback(auditCallback),
		jira.WithAutoRenewalToken(existingToken),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Client created with file-based token storage and audit logging")
	
	// Use the client - tokens will be automatically stored to file when refreshed
	myself, _, err := client.MySelf.Details(context.Background(), nil)
	if err != nil {
		log.Printf("Error getting user details: %v", err)
	} else {
		fmt.Printf("Authenticated as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)
	}
	
	// The token is now persisted to disk and can be loaded by other processes
	fmt.Printf("\nToken files stored in: %s\n", tokenStorePath)
	fmt.Printf("Audit log available at: %s\n", auditLog)
}

// Example: Simple usage without external dependencies
func main() {
	// This example shows the simplest possible usage
	
	// OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret", 
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Your existing token (from previous OAuth flow)
	token := &common.OAuth2Token{
		AccessToken:  "your-access-token",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		RefreshToken: "your-refresh-token",
		Scope:        "read:jira-work write:jira-work offline_access",
	}
	
	// Option 1: Basic auto-renewal (no external storage)
	client1, _ := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithAutoRenewalToken(token),
	)
	
	// Option 2: With file storage (tokens persist across restarts)
	tokenStore, _ := NewFileTokenStore("./.tokens")
	client2, _ := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithTokenStore(tokenStore),
		jira.WithAutoRenewalToken(token),
	)
	
	// Option 3: With callback for monitoring
	callback := &LoggingCallback{}
	client3, _ := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithTokenCallback(callback),
		jira.WithAutoRenewalToken(token),
	)
	
	fmt.Println("Clients created with different configurations")
	_ = client1
	_ = client2
	_ = client3
}

// LoggingCallback is a simple callback that logs token refresh events
type LoggingCallback struct{}

func (l *LoggingCallback) OnTokenRefreshed(ctx context.Context, oldToken, newToken *common.OAuth2Token) error {
	log.Printf("Token refreshed! New token expires in %d seconds", newToken.ExpiresIn)
	return nil
}