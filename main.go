package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
)

func renderPage(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from the query string
	urlToRender := r.URL.Query().Get("url")
	if urlToRender == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Create a context with a timeout to avoid hanging requests
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Timeout for the context
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(urlToRender),
		chromedp.Sleep(1*time.Second),    // Wait for JS to render. Adjust time as needed.
		chromedp.OuterHTML("html", &res), // Capture the outer HTML
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to render the page: %v", err), http.StatusInternalServerError)
		return
	}

	// Write the result to the response
	w.Write([]byte(res))
}

func main() {
	http.HandleFunc("/render", renderPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
