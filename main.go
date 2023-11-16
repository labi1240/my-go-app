package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const browserlessAPIKey = "487c6b40-dc79-4bdb-9bbe-c5d12064395b" // Replace with your actual API key

func renderHandler(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from the query string
	urlToRender := r.URL.Query().Get("url")
	if urlToRender == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Construct the Browserless API URL
	apiURL := fmt.Sprintf("https://chrome.browserless.io/content?token=%s&url=%s", browserlessAPIKey, url.QueryEscape(urlToRender))

	// Make the HTTP request to Browserless
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to make request to Browserless: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Browserless API returned status: %s", resp.Status), http.StatusInternalServerError)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read response from Browserless: %v", err), http.StatusInternalServerError)
		return
	}

	// Write the rendered HTML to the response
	w.Write(body)
}

func main() {
	http.HandleFunc("/render", renderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
