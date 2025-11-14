package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	// creating a new request
	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// setting a header
	req.Header.Set("User-Agent", "gator")

	// making the request
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()

	// decoding the xml to go structs
	rssFeed := RSSFeed{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rssFeed)
	if err != nil {
		return nil, err
	}

	// cleaing up the using html.UnescapeString
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	return &rssFeed, nil
}
