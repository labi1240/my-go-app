package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Set your Browserless API key and the URL to render
	apiKey := "487c6b40-dc79-4bdb-9bbe-c5d12064395b"
	targetURL := "https://www.example.com"

	// Construct the Browserless API URL
	apiURL := "https://chrome.browserless.io/content?token=" + apiKey + "&url=" + url.QueryEscape(targetURL)

	// Make the HTTP request to Browserless
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response: %v", err)
	}

	// Output the rendered HTML
	log.Printf("Rendered HTML: %s", body)
}
