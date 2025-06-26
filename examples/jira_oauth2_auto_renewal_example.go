package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// Example1BasicAutoRenewal demonstrates the simplest way to use OAuth 2.0 with automatic token renewal
func Example1BasicAutoRenewal() {
	// OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Your existing token (from previous OAuth flow)
	existingToken := &common.OAuth2Token{
		AccessToken:  "existing-access-token",
		TokenType:    "Bearer",
		ExpiresIn:    3600, // 1 hour
		RefreshToken: "existing-refresh-token",
		Scope:        "read:jira-work write:jira-work offline_access",
	}
	
	// Create client with auto-renewal - tokens will be refreshed automatically
	client, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithAutoRenewalToken(existingToken),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Use the client normally - no manual token management needed
	myself, _, err := client.MySelf.Details(context.Background(), nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	
	fmt.Printf("Authenticated as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)
}

// Example2WithCustomStorage demonstrates using auto-renewal with custom token storage
func Example2WithCustomStorage() {
	// Simple in-memory storage implementation
	storage := &InMemoryTokenStore{
		tokens: make(map[string]interface{}),
	}
	
	// OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Existing token
	existingToken := &common.OAuth2Token{
		AccessToken:  "existing-access-token",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		RefreshToken: "existing-refresh-token",
		Scope:        "read:jira-work write:jira-work offline_access",
	}
	
	// Create client with custom storage
	client, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithTokenStore(storage),
		jira.WithAutoRenewalToken(existingToken),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Client created with custom token storage")
	
	// Use the client - tokens will be automatically stored when refreshed
	myself, _, err := client.MySelf.Details(context.Background(), nil)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Authenticated as: %s\n", myself.DisplayName)
	}
}

// InMemoryTokenStore is a simple in-memory implementation of oauth2.TokenStore
type InMemoryTokenStore struct {
	tokens map[string]interface{}
	mu     sync.RWMutex
}

func (s *InMemoryTokenStore) GetToken(ctx context.Context) (*common.OAuth2Token, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	if token, ok := s.tokens["token"].(*common.OAuth2Token); ok {
		return token, nil
	}
	return nil, fmt.Errorf("token not found")
}

func (s *InMemoryTokenStore) SetToken(ctx context.Context, token *common.OAuth2Token) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.tokens["token"] = token
	fmt.Printf("Token stored: expires in %d seconds\n", token.ExpiresIn)
	return nil
}

func (s *InMemoryTokenStore) GetRefreshToken(ctx context.Context) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	if refreshToken, ok := s.tokens["refresh_token"].(string); ok {
		return refreshToken, nil
	}
	return "", fmt.Errorf("refresh token not found")
}

func (s *InMemoryTokenStore) SetRefreshToken(ctx context.Context, refreshToken string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.tokens["refresh_token"] = refreshToken
	fmt.Println("Refresh token updated")
	return nil
}

// Example3MultiSiteAutoRenewal demonstrates using auto-renewal with multiple Atlassian sites
func Example3MultiSiteAutoRenewal() {
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Initial token (would be loaded from storage)
	token := &common.OAuth2Token{
		AccessToken:  "access-token",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		RefreshToken: "refresh-token",
		Scope:        "read:jira-work write:jira-work offline_access",
	}
	
	// Create a temporary client to get accessible resources
	tempClient, err := jira.New(http.DefaultClient, "https://api.atlassian.com", jira.WithOAuth(oauthConfig))
	if err != nil {
		log.Fatal(err)
	}
	
	// Get accessible resources
	resources, err := tempClient.OAuth.GetAccessibleResources(context.Background(), token.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	
	// Create clients for each site with auto-renewal
	clients := make(map[string]*jira.Client)
	for _, resource := range resources {
		client, err := jira.New(
			http.DefaultClient,
			resource.URL,
			jira.WithOAuth(oauthConfig),
			jira.WithAutoRenewalToken(token),
		)
		if err != nil {
			log.Printf("Failed to create client for %s: %v", resource.Name, err)
			continue
		}
		
		clients[resource.Name] = client
		fmt.Printf("Created auto-renewing client for site: %s\n", resource.Name)
	}
	
	// Use the clients - each will auto-renew tokens as needed
	for siteName, client := range clients {
		projects, _, err := client.Project.Search(context.Background(), nil, 0, 10)
		if err != nil {
			log.Printf("Error getting projects for %s: %v", siteName, err)
			continue
		}
		
		fmt.Printf("\nProjects in %s:\n", siteName)
		for _, project := range projects.Values {
			fmt.Printf("- %s (%s)\n", project.Name, project.Key)
		}
	}
}