//go:build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	jira "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/oauth2"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// RedisTokenStore implements oauth2.TokenStore using Redis
type RedisTokenStore struct {
	client    *redis.Client
	keyPrefix string
	ctx       context.Context
}

// NewRedisTokenStore creates a new Redis-based token store
func NewRedisTokenStore(client *redis.Client, keyPrefix string) *RedisTokenStore {
	return &RedisTokenStore{
		client:    client,
		keyPrefix: keyPrefix,
		ctx:       context.Background(),
	}
}

// GetToken retrieves the OAuth2 token from Redis
func (r *RedisTokenStore) GetToken(ctx context.Context) (*common.OAuth2Token, error) {
	key := r.keyPrefix + "token"
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("token not found in Redis")
		}
		return nil, fmt.Errorf("failed to get token from Redis: %w", err)
	}
	
	var token common.OAuth2Token
	if err := json.Unmarshal([]byte(data), &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}
	
	return &token, nil
}

// SetToken stores the OAuth2 token in Redis with expiration
func (r *RedisTokenStore) SetToken(ctx context.Context, token *common.OAuth2Token) error {
	data, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}
	
	key := r.keyPrefix + "token"
	// Set expiration slightly longer than token expiry to handle edge cases
	expiration := time.Duration(token.ExpiresIn+300) * time.Second
	
	if err := r.client.Set(ctx, key, data, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set token in Redis: %w", err)
	}
	
	return nil
}

// GetRefreshToken retrieves only the refresh token from Redis
func (r *RedisTokenStore) GetRefreshToken(ctx context.Context) (string, error) {
	key := r.keyPrefix + "refresh_token"
	refreshToken, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("refresh token not found in Redis")
		}
		return "", fmt.Errorf("failed to get refresh token from Redis: %w", err)
	}
	
	return refreshToken, nil
}

// SetRefreshToken stores only the refresh token in Redis
// IMPORTANT: This method MUST be reliable. Errors from this method will cause
// the entire token refresh operation to fail. This is by design - refresh tokens
// are critical and losing them means the user must re-authenticate.
func (r *RedisTokenStore) SetRefreshToken(ctx context.Context, refreshToken string) error {
	key := r.keyPrefix + "refresh_token"
	// Refresh tokens typically don't expire, but set a long TTL for safety
	expiration := 90 * 24 * time.Hour // 90 days
	
	if err := r.client.Set(ctx, key, refreshToken, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set refresh token in Redis: %w", err)
	}
	
	return nil
}

// LoggingTokenCallback implements oauth2.TokenCallback for logging token refresh events
type LoggingTokenCallback struct{}

func (l *LoggingTokenCallback) OnTokenRefreshed(ctx context.Context, oldToken, newToken *common.OAuth2Token) error {
	log.Printf("Token refreshed successfully")
	log.Printf("New token expires in: %d seconds", newToken.ExpiresIn)
	if oldToken != nil {
		log.Printf("Old token was about to expire")
	}
	return nil
}

// Example: Using OAuth 2.0 with Redis storage for distributed applications
func ExampleOAuth2WithRedisStorage() {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // default DB
	})
	
	// Test Redis connection
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	
	// Create Redis token store
	tokenStore := NewRedisTokenStore(redisClient, "atlassian:oauth:")
	
	// OAuth configuration
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	// Try to load existing token from Redis
	var existingToken *common.OAuth2Token
	storedToken, err := tokenStore.GetToken(context.Background())
	if err == nil {
		existingToken = storedToken
		fmt.Println("Loaded existing token from Redis")
	} else {
		// If no token in Redis, use a new one (in real app, this would come from OAuth flow)
		existingToken = &common.OAuth2Token{
			AccessToken:  "existing-access-token",
			TokenType:    "Bearer",
			ExpiresIn:    3600, // 1 hour
			RefreshToken: "existing-refresh-token",
			Scope:        "read:jira-work write:jira-work offline_access",
		}
		fmt.Println("Using new token (not found in Redis)")
	}
	
	// Create client with Redis storage and logging callback
	client, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithTokenStore(tokenStore),
		jira.WithTokenCallback(&LoggingTokenCallback{}),
		jira.WithAutoRenewalToken(existingToken),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Client created with Redis token storage")
	
	// Use the client - tokens will be automatically stored in Redis when refreshed
	myself, _, err := client.MySelf.Details(context.Background(), nil)
	if err != nil {
		log.Printf("Error getting user details: %v", err)
	} else {
		fmt.Printf("Authenticated as: %s (%s)\n", myself.DisplayName, myself.EmailAddress)
	}
	
	// Simulate multiple application instances
	fmt.Println("\nSimulating second application instance...")
	
	// Create another client instance (simulating another server/pod)
	// It will automatically load the token from Redis
	client2, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithTokenStore(tokenStore),
		jira.WithAutoRenewalToken(&common.OAuth2Token{
			// Provide minimal token - the real one will be loaded from Redis
			RefreshToken: "existing-refresh-token",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	// This will use the token from Redis (shared across instances)
	projects, _, err := client2.Project.Search(context.Background(), nil, 0, 5)
	if err != nil {
		log.Printf("Error getting projects: %v", err)
	} else {
		fmt.Println("\nProjects accessible by second instance:")
		for _, project := range projects.Values {
			fmt.Printf("- %s (%s)\n", project.Name, project.Key)
		}
	}
}

// Example: Using multiple callbacks
func ExampleMultipleCallbacks() {
	// Create multiple callbacks
	loggingCallback := &LoggingTokenCallback{}
	
	// Metrics callback
	metricsCallback := &MetricsTokenCallback{}
	
	// Combine callbacks
	compositeCallback := oauth2.NewCompositeTokenCallback(
		loggingCallback,
		metricsCallback,
	)
	
	// Use with client
	redisStore := NewRedisTokenStore(redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}), "atlassian:")
	
	oauthConfig := &common.OAuth2Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		RedirectURI:  "https://your-app.com/callback",
	}
	
	client, err := jira.New(
		http.DefaultClient,
		"https://your-domain.atlassian.net",
		jira.WithOAuth(oauthConfig),
		jira.WithTokenStore(redisStore),
		jira.WithTokenCallback(compositeCallback),
		jira.WithAutoRenewalToken(&common.OAuth2Token{
			AccessToken:  "token",
			RefreshToken: "refresh-token",
			ExpiresIn:    3600,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Client created with multiple callbacks")
	_ = client
}

// MetricsTokenCallback tracks token refresh metrics
type MetricsTokenCallback struct {
	refreshCount int
}

func (m *MetricsTokenCallback) OnTokenRefreshed(ctx context.Context, oldToken, newToken *common.OAuth2Token) error {
	m.refreshCount++
	log.Printf("Token refresh count: %d", m.refreshCount)
	// In a real app, you might send this to Prometheus, DataDog, etc.
	return nil
}