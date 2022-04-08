package main

import (
	"kubectl-breakingnews/pkg/news"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	countryCode := os.Getenv("COUNTRY_CODE")

	news.IndianBreakingLiveNews(apiKey, countryCode)
}
