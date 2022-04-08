package news

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type NewsStore struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      string      `json:"author"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		URL         string      `json:"url"`
		URLToImage  string      `json:"urlToImage"`
		PublishedAt time.Time   `json:"publishedAt"`
		Content     interface{} `json:"content"`
	} `json:"articles"`
}

func IndianBreakingLiveNews(apiKey string, countryCode string) {
	var news NewsStore
	newGetURL := "https://newsapi.org/v2/top-headlines?country=" + countryCode
	httpClient := http.Client{}
	newsRequest, err := http.NewRequest("GET", newGetURL, nil)
	if err != nil {
		log.Print("Unabe to create news request because of: ", err)
	}
	newsRequest.Header.Set("x-api-key", apiKey)
	newsResponse, err := httpClient.Do(newsRequest)
	if err != nil {
		log.Print("Unabe to get the news because of: ", err)
	}
	defer newsResponse.Body.Close()
	err = json.NewDecoder(newsResponse.Body).Decode(&news)
	if err != nil {
		log.Print("Unabe decode the news response: ", err)
	}
	for startNews := 0; startNews < len(news.Articles); startNews++ {
		fmt.Print("\t\n\n", "Breaking News: ", news.Articles[startNews].Title)
	}
}
