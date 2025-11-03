//go:build ignore

package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// OAuthServer handles the OAuth 2.0 flow with HTTP callbacks
type OAuthServer struct {
	config    *common.OAuth2Config
	client    *jira.Client
	server    *http.Server
	state     string
	token     *common.OAuth2Token
	mu        sync.Mutex
	completed chan bool
}

// NewOAuthServer creates a new OAuth server
func NewOAuthServer(config *common.OAuth2Config) *OAuthServer {
	// Generate a random state parameter for CSRF protection
	state := generateRandomState()
	
	// Create temporary client for OAuth flow
	client, err := jira.New(
		http.DefaultClient,
		"https://api.atlassian.com",
		jira.WithOAuth(config),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &OAuthServer{
		config:    config,
		client:    client,
		state:     state,
		completed: make(chan bool, 1),
	}
}

// Start begins the OAuth flow and starts the HTTP server
func (s *OAuthServer) Start() error {
	// Set up HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleHome)
	mux.HandleFunc("/login", s.handleLogin)
	mux.HandleFunc("/callback", s.handleCallback)
	mux.HandleFunc("/sites", s.handleSites)

	// Create server
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("OAuth server starting on http://localhost:8080")
	fmt.Println("Visit http://localhost:8080 to begin OAuth flow")

	return s.server.ListenAndServe()
}

// Stop gracefully shuts down the server
func (s *OAuthServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// WaitForCompletion waits for the OAuth flow to complete
func (s *OAuthServer) WaitForCompletion() *common.OAuth2Token {
	<-s.completed
	return s.token
}

// handleHome displays the main page
func (s *OAuthServer) handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := `<!DOCTYPE html>
<html>
<head>
	<title>Go-Atlassian OAuth Example</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.container { max-width: 600px; margin: 0 auto; }
		.btn { background: #0052cc; color: white; padding: 10px 20px; 
			   text-decoration: none; border-radius: 4px; display: inline-block; }
		.btn:hover { background: #003d99; }
		.status { margin: 20px 0; padding: 10px; border-radius: 4px; }
		.success { background: #e3fcef; border: 1px solid #00875a; color: #00875a; }
		.info { background: #deebff; border: 1px solid #0052cc; color: #0052cc; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Go-Atlassian OAuth 2.0 Example</h1>
		<p>This example demonstrates OAuth 2.0 authentication with Atlassian services.</p>
		
		{{if .Token}}
		<div class="status success">
			<strong>✅ OAuth flow completed successfully!</strong><br>
			Access token obtained and ready to use.
		</div>
		<p><a href="/sites" class="btn">View Accessible Sites</a></p>
		{{else}}
		<div class="status info">
			<strong>ℹ️ Ready to authenticate</strong><br>
			Click the button below to start the OAuth flow.
		</div>
		<p><a href="/login" class="btn">Login with Atlassian</a></p>
		{{end}}
		
		<h3>OAuth Configuration</h3>
		<ul>
			<li><strong>Client ID:</strong> {{.Config.ClientID}}</li>
			<li><strong>Redirect URI:</strong> {{.Config.RedirectURI}}</li>
			<li><strong>Scopes:</strong> {{range .Config.Scopes}}{{.}} {{end}}</li>
		</ul>
	</div>
</body>
</html>`

	t := template.Must(template.New("home").Parse(tmpl))
	data := struct {
		Config *common.OAuth2Config
		Token  *common.OAuth2Token
	}{
		Config: s.config,
		Token:  s.token,
	}
	
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handleLogin redirects to Atlassian OAuth authorization
func (s *OAuthServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	scopes := []string{"read:jira-work", "write:jira-work", "read:jira-user"}
	authURL, err := s.client.OAuth.GetAuthorizationURL(scopes, s.state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate auth URL: %v", err), 
			http.StatusInternalServerError)
		return
	}

	log.Printf("Redirecting to authorization URL: %s", authURL.String())
	http.Redirect(w, r, authURL.String(), http.StatusFound)
}

// handleCallback processes the OAuth callback
func (s *OAuthServer) handleCallback(w http.ResponseWriter, r *http.Request) {
	// Extract parameters from callback URL
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	errorParam := r.URL.Query().Get("error")
	errorDescription := r.URL.Query().Get("error_description")

	// Check for OAuth errors
	if errorParam != "" {
		s.renderError(w, fmt.Sprintf("OAuth error: %s - %s", 
			errorParam, errorDescription))
		return
	}

	// Verify state parameter (CSRF protection)
	if state != s.state {
		s.renderError(w, "Invalid state parameter - possible CSRF attack")
		return
	}

	// Ensure we have an authorization code
	if code == "" {
		s.renderError(w, "No authorization code received")
		return
	}

	log.Printf("Received authorization code: %s", code[:10]+"...")

	// Exchange authorization code for tokens
	ctx := context.Background()
	token, err := s.client.OAuth.ExchangeAuthorizationCode(ctx, code)
	if err != nil {
		s.renderError(w, fmt.Sprintf("Failed to exchange code for token: %v", err))
		return
	}

	// Store token
	s.mu.Lock()
	s.token = token
	s.mu.Unlock()

	log.Printf("OAuth flow completed successfully!")
	log.Printf("Access token: %s...", token.AccessToken[:20])
	log.Printf("Token expires in: %d seconds", token.ExpiresIn)

	// Signal completion
	select {
	case s.completed <- true:
	default:
	}

	// Redirect to home page
	http.Redirect(w, r, "/", http.StatusFound)
}

// handleSites displays accessible Atlassian sites
func (s *OAuthServer) handleSites(w http.ResponseWriter, r *http.Request) {
	if s.token == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Get accessible resources
	ctx := context.Background()
	resources, err := s.client.OAuth.GetAccessibleResources(ctx, s.token.AccessToken)
	if err != nil {
		s.renderError(w, fmt.Sprintf("Failed to get accessible resources: %v", err))
		return
	}

	tmpl := `<!DOCTYPE html>
<html>
<head>
	<title>Accessible Sites - Go-Atlassian OAuth</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.container { max-width: 800px; margin: 0 auto; }
		.site { border: 1px solid #ddd; margin: 10px 0; padding: 15px; border-radius: 4px; }
		.site h3 { margin: 0 0 10px 0; color: #0052cc; }
		.btn { background: #0052cc; color: white; padding: 8px 16px; 
			   text-decoration: none; border-radius: 4px; display: inline-block; margin-top: 10px; }
		.btn:hover { background: #003d99; }
		.scopes { color: #666; font-size: 0.9em; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Accessible Atlassian Sites</h1>
		<p>Your OAuth token provides access to the following sites:</p>
		
		{{range .Resources}}
		<div class="site">
			<h3>{{.Name}}</h3>
			<p><strong>URL:</strong> <a href="{{.URL}}" target="_blank">{{.URL}}</a></p>
			<p><strong>Site ID:</strong> {{.ID}}</p>
			<div class="scopes">
				<strong>Available scopes:</strong> {{range .Scopes}}{{.}} {{end}}
			</div>
		</div>
		{{else}}
		<p>No accessible sites found.</p>
		{{end}}
		
		<p><a href="/" class="btn">← Back to Home</a></p>
	</div>
</body>
</html>`

	t := template.Must(template.New("sites").Parse(tmpl))
	data := struct {
		Resources []*common.AccessibleResource
	}{
		Resources: resources,
	}
	
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// renderError displays an error page
func (s *OAuthServer) renderError(w http.ResponseWriter, message string) {
	tmpl := `<!DOCTYPE html>
<html>
<head>
	<title>Error - Go-Atlassian OAuth</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.container { max-width: 600px; margin: 0 auto; }
		.error { background: #ffebe6; border: 1px solid #de350b; color: #de350b; 
				 padding: 15px; border-radius: 4px; margin: 20px 0; }
		.btn { background: #0052cc; color: white; padding: 10px 20px; 
			   text-decoration: none; border-radius: 4px; display: inline-block; }
	</style>
</head>
<body>
	<div class="container">
		<h1>OAuth Error</h1>
		<div class="error">
			<strong>❌ Error:</strong> {{.}}
		</div>
		<p><a href="/" class="btn">← Back to Home</a></p>
	</div>
</body>
</html>`

	w.WriteHeader(http.StatusBadRequest)
	t := template.Must(template.New("error").Parse(tmpl))
	if err := t.Execute(w, message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// generateRandomState creates a random state parameter for CSRF protection
func generateRandomState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// ExampleHTTPServerOAuth demonstrates OAuth with HTTP server
func ExampleHTTPServerOAuth() {
	// OAuth configuration
	config := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "http://localhost:8080/callback",
		Scopes: []string{
			"read:jira-work",
			"write:jira-work",
			"read:jira-user",
		},
	}

	// Create OAuth server
	server := NewOAuthServer(config)

	// Start server in a goroutine
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for OAuth completion or timeout
	select {
	case <-time.After(5 * time.Minute):
		log.Println("OAuth flow timed out")
		if err := server.Stop(); err != nil {
			log.Printf("Error stopping server: %v", err)
		}
		return
	case token := <-server.completed:
		if token {
			log.Println("OAuth flow completed successfully!")
			
			// Get the token
			oauthToken := server.WaitForCompletion()
			
			// Now create a client with auto-renewal for actual API usage
			// First, get accessible resources to determine site URL
			resources, err := server.client.OAuth.GetAccessibleResources(
				context.Background(), 
				oauthToken.AccessToken,
			)
			if err != nil {
				log.Printf("Failed to get resources: %v", err)
				if stopErr := server.Stop(); stopErr != nil {
					log.Printf("Error stopping server: %v", stopErr)
				}
				return
			}

			if len(resources) == 0 {
				log.Println("No accessible resources found")
				if stopErr := server.Stop(); stopErr != nil {
					log.Printf("Error stopping server: %v", stopErr)
				}
				return
			}

			// Create production client with auto-renewal
			client, err := jira.New(
				http.DefaultClient,
				resources[0].URL,
				jira.WithOAuth(config),
				jira.WithAutoRenewalToken(oauthToken),
			)
			if err != nil {
				log.Printf("Failed to create client: %v", err)
				if stopErr := server.Stop(); stopErr != nil {
					log.Printf("Error stopping server: %v", stopErr)
				}
				return
			}

			// Test the API
			myself, _, err := client.MySelf.Details(context.Background(), nil)
			if err != nil {
				log.Printf("API call failed: %v", err)
			} else {
				log.Printf("Successfully authenticated as: %s (%s)", 
					myself.DisplayName, myself.EmailAddress)
			}
		}
	}

	// Keep server running for demonstration
	// In a real application, you might want to handle shutdown signals
	log.Println("Server will continue running. Press Ctrl+C to stop.")
	select {} // Block forever
}

func main() {
	ExampleHTTPServerOAuth()
}