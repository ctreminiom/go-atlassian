package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// Global variables for this simple example
var (
	oauthConfig *common.OAuth2Config
	oauthClient *jira.Client
	tokenChan   = make(chan *common.OAuth2Token, 1)
	state       = "unique-random-state" // In production, generate this randomly
)

// ExampleSimpleHTTPServerOAuth demonstrates a minimal OAuth flow with HTTP server
func ExampleSimpleHTTPServerOAuth() {
	// Step 1: Configure OAuth
	oauthConfig = &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "http://localhost:8080/callback",
	}

	// Step 2: Create OAuth client
	var err error
	oauthClient, err = jira.New(
		http.DefaultClient,
		"https://api.atlassian.com",
		jira.WithOAuth(oauthConfig),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Step 3: Set up HTTP routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	// Step 4: Start server
	fmt.Println("OAuth server starting on http://localhost:8080")
	fmt.Println("Visit http://localhost:8080/login to start OAuth flow")

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Step 5: Wait for OAuth completion
	select {
	case token := <-tokenChan:
		fmt.Println("✅ OAuth completed successfully!")
		
		// Step 6: Use the token with auto-renewal
		useTokenWithAutoRenewal(token)
		
	case <-time.After(5 * time.Minute):
		fmt.Println("❌ OAuth flow timed out")
	}
}

// handleHome shows a simple home page
func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head><title>Simple OAuth Example</title></head>
<body>
	<h1>Go-Atlassian OAuth Example</h1>
	<p><a href="/login">Click here to login with Atlassian</a></p>
</body>
</html>`
	fmt.Fprint(w, html)
}

// handleLogin redirects to Atlassian OAuth
func handleLogin(w http.ResponseWriter, r *http.Request) {
	scopes := []string{"read:jira-work", "write:jira-work"}
	authURL, err := oauthClient.OAuth.GetAuthorizationURL(scopes, state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Redirecting to: %s\n", authURL.String())
	http.Redirect(w, r, authURL.String(), http.StatusFound)
}

// handleCallback processes the OAuth callback
func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Extract parameters
	code := r.URL.Query().Get("code")
	receivedState := r.URL.Query().Get("state")
	errorParam := r.URL.Query().Get("error")

	// Check for errors
	if errorParam != "" {
		errorDesc := r.URL.Query().Get("error_description")
		message := fmt.Sprintf("OAuth error: %s - %s", errorParam, errorDesc)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	// Verify state (CSRF protection)
	if receivedState != state {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	// Exchange code for token
	ctx := context.Background()
	token, err := oauthClient.OAuth.ExchangeAuthorizationCode(ctx, code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Token exchange failed: %v", err), 
			http.StatusInternalServerError)
		return
	}

	// Send token to main goroutine
	select {
	case tokenChan <- token:
		// Success response
		html := `
<!DOCTYPE html>
<html>
<head><title>OAuth Success</title></head>
<body>
	<h1>✅ OAuth Successful!</h1>
	<p>You have been successfully authenticated with Atlassian.</p>
	<p>You can now close this window.</p>
</body>
</html>`
		fmt.Fprint(w, html)
	default:
		http.Error(w, "Token channel full", http.StatusInternalServerError)
	}
}

// useTokenWithAutoRenewal demonstrates using the token with auto-renewal
func useTokenWithAutoRenewal(token *common.OAuth2Token) {
	fmt.Printf("Access token: %s...\n", token.AccessToken[:20])
	fmt.Printf("Expires in: %d seconds\n", token.ExpiresIn)

	// Get accessible resources first
	ctx := context.Background()
	resources, err := oauthClient.OAuth.GetAccessibleResources(ctx, token.AccessToken)
	if err != nil {
		log.Printf("Failed to get accessible resources: %v", err)
		return
	}

	if len(resources) == 0 {
		fmt.Println("No accessible Atlassian sites found")
		return
	}

	// Create client with auto-renewal for the first site
	client, err := jira.New(
		http.DefaultClient,
		resources[0].URL,
		jira.WithOAuth(oauthConfig),
		jira.WithAutoRenewalToken(token),
	)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return
	}

	fmt.Printf("Connected to site: %s (%s)\n", resources[0].Name, resources[0].URL)

	// Test API call
	myself, _, err := client.MySelf.Details(ctx, nil)
	if err != nil {
		log.Printf("API call failed: %v", err)
		return
	}

	fmt.Printf("Authenticated as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)

	// Demonstrate that the client works with auto-renewal
	fmt.Println("Client is ready for use with automatic token renewal!")
	
	// Example: Get projects
	projects, _, err := client.Project.Search(ctx, nil, 0, 5)
	if err != nil {
		log.Printf("Failed to get projects: %v", err)
		return
	}

	fmt.Printf("Found %d projects:\n", len(projects.Values))
	for _, project := range projects.Values {
		fmt.Printf("- %s (%s)\n", project.Name, project.Key)
	}
}