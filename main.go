// main.go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dghubble/oauth1"
)

// Twitter API credentials
const (
	consumerKey    = "" // Replace with API Key
	consumerSecret = "" // Replace with API Secret Key
	accessToken    = "" // Replace with Access Token
	accessSecret   = "" // Replace with Access Token Secret
	delaySecond    = 30 //setting of delay second
)

func main() {
	// Example of posting a tweet
	tweetID, err := postTweet("Hello from Twitter API!")
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}
	fmt.Printf("Posted tweet with ID: %s\n", tweetID)
	fmt.Printf("Waiting for %d seconds before deleting a tweet...\n", delaySecond)
	time.Sleep(time.Duration(delaySecond) * time.Second)

	//Deleting the tweet after delay
	err = deleteTweet(tweetID)
	if err != nil {
		log.Fatalf("Error deleting tweet: %v", err)

	}
	fmt.Printf("Deleted tweet with ID: %s\n", tweetID)

}

// postTweet sends a tweet to Twitter
func postTweet(content string) (string, error) {
	// OAuth1 authentication
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create HTTP client with OAuth1
	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	// Create tweet content
	tweetData := map[string]interface{}{
		"text": content,
	}
	tweetJSON, err := json.Marshal(tweetData)
	if err != nil {
		return "", fmt.Errorf("error marshaling tweet content: %w", err)
	}

	// Make POST request to v2 tweets endpoint
	response, err := httpClient.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBuffer(tweetJSON))
	if err != nil {
		return "", fmt.Errorf("failed to post tweet: %w", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode == http.StatusCreated {
		var result map[string]interface{}
		body, _ := ioutil.ReadAll(response.Body)
		if err := json.Unmarshal(body, &result); err != nil {
			return "", fmt.Errorf("error unmarshaling response: %w", err)
		}
		// Extract tweet ID from the response
		return result["data"].(map[string]interface{})["id"].(string), nil
	}

	// Print detailed error information
	body, _ := ioutil.ReadAll(response.Body)
	return "", fmt.Errorf("failed to post tweet: %s, response: %s", response.Status, string(body))
}

func deleteTweet(tweetID string) error {
	// OAuth1 authentication
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Create HTTP client with OAuth1
	ctx := context.Background()
	httpClient := config.Client(ctx, token)

	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID), nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	// Execute the DELETE request
	response, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute delete request: %w", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode != http.StatusNoContent {
		// Print detailed error information
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("failed to delete tweet: %s, response: %s", response.Status, string(body))
	}

	return nil
}
