package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
)

func main() {
	// Example of using OAuth 2.0 with go-atlassian
	
	// Step 1: Create a client with OAuth configuration
	client, err := jira.NewWithOAuth(
		http.DefaultClient,
		"your-client-id",
		"your-client-secret",
		"https://your-app.com/callback",
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// Step 2: Generate authorization URL
	scopes := []string{
		"read:jira-work",
		"write:jira-work",
		"read:jira-user",
		"manage:jira-project",
	}
	
	authURL, err := client.OAuth.GetAuthorizationURL(scopes, "unique-state-value")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Visit this URL to authorize the application:\n%s\n", authURL.String())
	
	// Step 3: After user authorizes, they will be redirected to your callback URL
	// Extract the authorization code from the callback URL query parameters
	// For example: https://your-app.com/callback?code=AUTH_CODE&state=unique-state-value
	
	// Step 4: Exchange authorization code for tokens
	authCode := "authorization-code-from-callback"
	token, err := client.OAuth.ExchangeAuthorizationCode(context.Background(), authCode)
	if err != nil {
		log.Fatal(err)
	}
	
	// Store the access token and refresh token
	client.Auth.SetOAuth2AccessToken(token.AccessToken)
	client.Auth.SetOAuth2RefreshToken(token.RefreshToken)
	
	fmt.Printf("Access token obtained, expires in %d seconds\n", token.ExpiresIn)
	
	// Step 5: Get accessible resources (Atlassian sites)
	resources, err := client.OAuth.GetAccessibleResources(context.Background(), token.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	
	if len(resources) == 0 {
		log.Fatal("No accessible resources found")
	}
	
	// Step 6: Set the site URL for the first accessible resource
	firstSite := resources[0]
	err = client.SetSiteURL(firstSite.URL)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Using site: %s (%s)\n", firstSite.Name, firstSite.URL)
	
	// Step 7: Now you can use the client to make API calls
	myself, _, err := client.MySelf.GetCurrentUser(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Authenticated as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)
	
	// Example: Get all projects
	projects, _, err := client.Project.Gets(context.Background(), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Found %d projects\n", len(projects.Values))
	for _, project := range projects.Values {
		fmt.Printf("- %s (%s)\n", project.Name, project.Key)
	}
	
	// Step 8: Refresh token when needed
	// Typically, you would check if the token is expired before making API calls
	newToken, err := client.OAuth.RefreshAccessToken(context.Background(), token.RefreshToken)
	if err != nil {
		log.Fatal(err)
	}
	
	// Update the tokens
	client.Auth.SetOAuth2AccessToken(newToken.AccessToken)
	client.Auth.SetOAuth2RefreshToken(newToken.RefreshToken)
	
	fmt.Println("Token refreshed successfully")
}

// Example callback handler for your web server
func callbackHandler(w http.ResponseWriter, r *http.Request) {
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
}