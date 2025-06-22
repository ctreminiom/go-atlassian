package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// ExampleOAuth2NewFlow demonstrates how to use the OAuth flow with the new separated API
func ExampleOAuth2NewFlow() {
	// OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
		Scopes: []string{
			"read:jira-work",
			"write:jira-work",
			"read:jira-user",
		},
	}
	
	// Step 1: Create a client with OAuth support for the authorization flow
	tempClient, err := jira.New(http.DefaultClient, "https://api.atlassian.com", jira.WithOAuth(oauthConfig))
	if err != nil {
		log.Fatal(err)
	}
	
	// Step 2: Generate authorization URL
	authURL, err := tempClient.OAuth.GetAuthorizationURL(oauthConfig.Scopes, "unique-state-value")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Visit this URL to authorize: %s\n", authURL.String())
	
	// Step 3: After user authorizes, exchange code for token
	authCode := "authorization-code-from-callback"
	token, err := tempClient.OAuth.ExchangeAuthorizationCode(context.Background(), authCode)
	if err != nil {
		log.Fatal(err)
	}
	
	// Step 4: Get accessible resources
	resources, err := tempClient.OAuth.GetAccessibleResources(context.Background(), token.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	
	if len(resources) == 0 {
		log.Fatal("No accessible resources found")
	}
	
	// Step 5: Create the actual client with auto-renewal for the first site
	client, err := jira.New(
		http.DefaultClient,
		resources[0].URL,
		jira.WithOAuth(oauthConfig),          // OAuth service for token refresh
		jira.WithAutoRenewalToken(token),      // Enable auto-renewal
	)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Connected to site: %s (%s)\n", resources[0].Name, resources[0].URL)
	
	// Step 6: Use the client - tokens will be automatically renewed
	myself, _, err := client.MySelf.Details(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Authenticated as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)
}

// ExampleManualTokenManagement shows how to use OAuth without auto-renewal
func ExampleManualTokenManagement() {
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Create client with OAuth but without auto-renewal
	client, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Manually set the token
	token := &common.OAuth2Token{
		AccessToken:  "your-access-token",
		RefreshToken: "your-refresh-token",
		ExpiresIn:    3600,
	}
	client.Auth.SetBearerToken(token.AccessToken)
	
	// Later, manually refresh when needed
	newToken, err := client.OAuth.RefreshAccessToken(context.Background(), token.RefreshToken)
	if err != nil {
		log.Printf("Failed to refresh token: %v", err)
		return
	}
	
	// Update the token manually
	client.Auth.SetBearerToken(newToken.AccessToken)
	fmt.Println("Token refreshed manually")
}

// ExampleMigrationPath shows how to migrate from manual to auto-renewal
func ExampleMigrationPath() {
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Start with manual token management
	client, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Use manual token for a while
	token := &common.OAuth2Token{
		AccessToken:  "your-access-token",
		RefreshToken: "your-refresh-token",
		ExpiresIn:    3600,
	}
	client.Auth.SetBearerToken(token.AccessToken)
	
	// Later, decide to switch to auto-renewal
	// Create a new client with auto-renewal enabled
	autoRenewClient, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithAutoRenewalToken(token),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Now use autoRenewClient instead of client
	fmt.Println("Migrated to auto-renewal client")
	_ = autoRenewClient
}