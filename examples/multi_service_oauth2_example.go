package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	admin "github.com/ctreminiom/go-atlassian/v2/admin"
	confluence "github.com/ctreminiom/go-atlassian/v2/confluence/v2"
	jirav2 "github.com/ctreminiom/go-atlassian/v2/jira/v2"
	jirav3 "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// ExampleMultiServiceOAuth demonstrates using OAuth 2.0 across multiple Atlassian services
func ExampleMultiServiceOAuth() {
	// Step 1: OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}

	// Step 2: Assume you already have an OAuth token (from previous authorization)
	// In a real application, you would obtain this through the OAuth flow
	token := &common.OAuth2Token{
		AccessToken:  "your-access-token",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		RefreshToken: "your-refresh-token",
		Scope:        "read:jira-work write:jira-work read:confluence-content manage:jira-configuration",
	}

	fmt.Println("🚀 Creating multi-service clients with OAuth 2.0...")

	// Step 3: Get accessible resources to determine available sites
	tempClient, err := jirav3.New(
		http.DefaultClient,
		"https://api.atlassian.com",
		jirav3.WithOAuth(oauthConfig),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	resources, err := tempClient.OAuth.GetAccessibleResources(ctx, token.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	if len(resources) == 0 {
		log.Fatal("No accessible Atlassian sites found")
	}

	// Use the first available site
	siteURL := resources[0].URL
	fmt.Printf("🌐 Using site: %s (%s)\n", resources[0].Name, siteURL)

	// Step 4: Create clients for multiple services with auto-renewal

	// Jira v3 Client (modern ADF format)
	jiraV3Client, err := jirav3.New(
		http.DefaultClient,
		siteURL,
		jirav3.WithOAuth(oauthConfig),
		jirav3.WithAutoRenewalToken(token),
	)
	if err != nil {
		log.Printf("❌ Failed to create Jira v3 client: %v", err)
	} else {
		fmt.Println("✅ Jira v3 client created with auto-renewal")
	}

	// Jira v2 Client (legacy Rich Text format)
	jiraV2Client, err := jirav2.New(
		http.DefaultClient,
		siteURL,
		jirav2.WithOAuth(oauthConfig),
		jirav2.WithAutoRenewalToken(token),
	)
	if err != nil {
		log.Printf("❌ Failed to create Jira v2 client: %v", err)
	} else {
		fmt.Println("✅ Jira v2 client created with auto-renewal")
	}

	// Confluence v2 Client
	confluenceClient, err := confluence.New(
		http.DefaultClient,
		siteURL,
		confluence.WithOAuth(oauthConfig),
		confluence.WithAutoRenewalToken(token),
	)
	if err != nil {
		log.Printf("❌ Failed to create Confluence client: %v", err)
	} else {
		fmt.Println("✅ Confluence v2 client created with auto-renewal")
	}

	// Admin Client (for organization management)
	adminClient, err := admin.New(
		http.DefaultClient,
		admin.WithOAuth(oauthConfig),
		admin.WithAutoRenewalToken(token),
	)
	if err != nil {
		log.Printf("❌ Failed to create Admin client: %v", err)
	} else {
		fmt.Println("✅ Admin client created with auto-renewal")
	}

	fmt.Println("\n🔑 Testing authentication across all services...")

	// Step 5: Test each service with the same OAuth token

	// Test Jira v3
	if jiraV3Client != nil {
		myself, _, err := jiraV3Client.MySelf.Details(ctx, nil)
		if err != nil {
			log.Printf("❌ Jira v3 authentication failed: %v", err)
		} else {
			fmt.Printf("✅ Jira v3: Authenticated as %s (%s)\n", myself.DisplayName, myself.EmailAddress)
		}

		// Get recent projects
		projects, _, err := jiraV3Client.Project.Search(ctx, nil, 0, 3)
		if err != nil {
			log.Printf("❌ Jira v3 project search failed: %v", err)
		} else {
			fmt.Printf("   📁 Found %d projects in Jira\n", len(projects.Values))
		}
	}

	// Test Jira v2
	if jiraV2Client != nil {
		myself, _, err := jiraV2Client.MySelf.Details(ctx, nil)
		if err != nil {
			log.Printf("❌ Jira v2 authentication failed: %v", err)
		} else {
			fmt.Printf("✅ Jira v2: Authenticated as %s (%s)\n", myself.DisplayName, myself.EmailAddress)
		}
	}

	// Test Confluence
	if confluenceClient != nil {
		spaces, _, err := confluenceClient.Space.Gets(ctx, nil, 0, 3)
		if err != nil {
			log.Printf("❌ Confluence space list failed: %v", err)
		} else {
			fmt.Printf("✅ Confluence: Found %d spaces\n", len(spaces.Results))
			for _, space := range spaces.Results {
				fmt.Printf("   📚 Space: %s (%s)\n", space.Name, space.Key)
			}
		}
	}

	// Test Admin API
	if adminClient != nil {
		orgs, _, err := adminClient.Organization.Gets(ctx, nil)
		if err != nil {
			log.Printf("❌ Admin organization list failed: %v", err)
		} else {
			fmt.Printf("✅ Admin: Found %d organizations\n", len(orgs))
			for _, org := range orgs {
				fmt.Printf("   🏢 Organization: %s\n", org.Name)
			}
		}
	}

	fmt.Println("\n🎉 Multi-service OAuth demonstration completed!")
	fmt.Println("All clients are using the same OAuth token with automatic renewal.")
}

// ExampleCrossServiceWorkflow demonstrates a workflow that uses multiple services
func ExampleCrossServiceWorkflow() {
	// This example shows how you might use multiple services together
	// in a real application workflow

	fmt.Println("\n🔄 Cross-service workflow example...")

	// OAuth setup (same as above)
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}

	token := &common.OAuth2Token{
		AccessToken:  "your-access-token",
		RefreshToken: "your-refresh-token",
		ExpiresIn:    3600,
	}

	// Get site URL
	tempClient, _ := jirav3.New(http.DefaultClient, "https://api.atlassian.com", jirav3.WithOAuth(oauthConfig))
	resources, _ := tempClient.OAuth.GetAccessibleResources(context.Background(), token.AccessToken)
	if len(resources) == 0 {
		return
	}
	siteURL := resources[0].URL

	// Create clients
	jiraClient, _ := jirav3.New(http.DefaultClient, siteURL, jirav3.WithOAuth(oauthConfig), jirav3.WithAutoRenewalToken(token))
	confluenceClient, _ := confluence.New(http.DefaultClient, siteURL, confluence.WithOAuth(oauthConfig), confluence.WithAutoRenewalToken(token))
	adminClient, _ := admin.New(http.DefaultClient, admin.WithOAuth(oauthConfig), admin.WithAutoRenewalToken(token))

	ctx := context.Background()

	// Workflow: User management across services
	fmt.Println("👤 User Management Workflow:")

	// 1. Get user info from Admin API
	if adminClient != nil {
		fmt.Println("   1. Getting user information from Admin API...")
		// In a real app, you'd get specific user details
	}

	// 2. Check user's Jira projects
	if jiraClient != nil {
		fmt.Println("   2. Checking user's Jira projects...")
		projects, _, err := jiraClient.Project.Search(ctx, nil, 0, 5)
		if err == nil {
			fmt.Printf("      Found %d accessible projects\n", len(projects.Values))
		}
	}

	// 3. Check user's Confluence spaces
	if confluenceClient != nil {
		fmt.Println("   3. Checking user's Confluence spaces...")
		spaces, _, err := confluenceClient.Space.Gets(ctx, nil, 0, 5)
		if err == nil {
			fmt.Printf("      Found %d accessible spaces\n", len(spaces.Results))
		}
	}

	fmt.Println("   ✅ Cross-service workflow completed!")
}

// ExampleServiceSpecificScopes demonstrates requesting service-specific scopes
func ExampleServiceSpecificScopes() {
	fmt.Println("\n🎯 Service-specific OAuth scopes example...")

	// Different OAuth configs for different use cases
	configs := map[string]*common.OAuth2Config{
		"jira-only": {
			ClientID:     "your-client-id",
			ClientSecret: "your-client-secret",
			RedirectURI:  "https://your-app.com/callback",
			Scopes: []string{
				"read:jira-work",
				"write:jira-work",
				"read:jira-user",
				"manage:jira-project",
			},
		},
		"confluence-only": {
			ClientID:     "your-client-id",
			ClientSecret: "your-client-secret",
			RedirectURI:  "https://your-app.com/callback",
			Scopes: []string{
				"read:confluence-content.summary",
				"read:confluence-content.all",
				"write:confluence-content",
				"read:confluence-space.summary",
			},
		},
		"admin-only": {
			ClientID:     "your-client-id",
			ClientSecret: "your-client-secret",
			RedirectURI:  "https://your-app.com/callback",
			Scopes: []string{
				"manage:jira-configuration",
				"read:account",
				"manage:organization",
			},
		},
		"multi-service": {
			ClientID:     "your-client-id",
			ClientSecret: "your-client-secret",
			RedirectURI:  "https://your-app.com/callback",
			Scopes: []string{
				"read:jira-work",
				"write:jira-work",
				"read:confluence-content.all",
				"write:confluence-content",
				"manage:jira-configuration",
				"read:account",
			},
		},
	}

	for name, config := range configs {
		fmt.Printf("📋 %s scopes: %v\n", name, config.Scopes)
	}

	fmt.Println("\nChoose the appropriate scope set based on your application's needs!")
}

func main() {
	ExampleMultiServiceOAuth()
	ExampleCrossServiceWorkflow()
	ExampleServiceSpecificScopes()
}