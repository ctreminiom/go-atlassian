package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

func main() {
	// Example of using OAuth 2.0 with go-atlassian
	
	// Step 1: Prepare OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
		Scopes: []string{
			"read:jira-work",
			"write:jira-work",
			"read:jira-user",
			"manage:jira-project",
		},
	}
	
	// Step 2: Create a client with OAuth support
	// Note: For OAuth flow, we first create a client with empty site URL
	client, err := jira.New(http.DefaultClient, "", jira.WithOAuth(oauthConfig))
	if err != nil {
		log.Fatal(err)
	}
	
	// Step 3: Generate authorization URL
	authURL, err := client.OAuth.GetAuthorizationURL(oauthConfig.Scopes, "unique-state-value")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Visit this URL to authorize the application:\n%s\n", authURL.String())
	
	// Step 4: After user authorizes, they will be redirected to your callback URL
	// Extract the authorization code from the callback URL query parameters
	// For example: https://your-app.com/callback?code=AUTH_CODE&state=unique-state-value
	
	// Step 5: Exchange authorization code for tokens
	authCode := "authorization-code-from-callback"
	token, err := client.OAuth.ExchangeAuthorizationCode(context.Background(), authCode)
	if err != nil {
		log.Fatal(err)
	}
	
	// Store the access token for API calls
	client.Auth.SetBearerToken(token.AccessToken)
	
	fmt.Printf("Access token obtained, expires in %d seconds\n", token.ExpiresIn)
	
	// Step 6: Get accessible resources (Atlassian sites)
	resources, err := client.OAuth.GetAccessibleResources(context.Background(), token.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	
	if len(resources) == 0 {
		log.Fatal("No accessible resources found")
	}
	
	// Step 7: Create a new client with the site URL
	firstSite := resources[0]
	client, err = jira.New(
		http.DefaultClient, 
		firstSite.URL,
		jira.WithOAuth(oauthConfig),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Set the OAuth token
	client.Auth.SetBearerToken(token.AccessToken)
	
	fmt.Printf("Using site: %s (%s)\n", firstSite.Name, firstSite.URL)
	
	// Step 8: Now you can use the client to make API calls
	myself, _, err := client.MySelf.Details(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Authenticated as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)
	
	// Example: Search for projects
	projects, _, err := client.Project.Search(context.Background(), nil, 0, 50)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Found %d projects\n", len(projects.Values))
	for _, project := range projects.Values {
		fmt.Printf("- %s (%s)\n", project.Name, project.Key)
	}
	
	// Step 9: Refresh token when needed
	// Typically, you would check if the token is expired before making API calls
	// and store the refresh token securely
	if token.RefreshToken != "" {
		newToken, err := client.OAuth.RefreshAccessToken(context.Background(), token.RefreshToken)
		if err != nil {
			log.Printf("Failed to refresh token: %v", err)
		} else {
			// Update the access token
			client.Auth.SetBearerToken(newToken.AccessToken)
			fmt.Println("Token refreshed successfully")
			
			// Store the new refresh token if provided
			if newToken.RefreshToken != "" {
				// Store this securely for future use
				token.RefreshToken = newToken.RefreshToken
			}
		}
	}
}

// ExampleCallbackHandler demonstrates how to handle the OAuth callback in your web server
func ExampleCallbackHandler() {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Extract authorization code and state from query parameters
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		error := r.URL.Query().Get("error")
		errorDescription := r.URL.Query().Get("error_description")
		
		if error != "" {
			fmt.Fprintf(w, "Authorization failed: %s - %s", error, errorDescription)
			return
		}
		
		// Verify state parameter matches what you sent
		// This prevents CSRF attacks
		if state != "unique-state-value" {
			fmt.Fprintf(w, "Invalid state parameter")
			return
		}
		
		// Use the authorization code to get tokens
		// ... (see main function for token exchange example)
		
		fmt.Fprintf(w, "Authorization successful! Code: %s", code)
	})
	
	// Example of starting the server (commented out to avoid actually running)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}