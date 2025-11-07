package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/anxhukumar/gator-cli-tool/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {

	// check if arguments are available
	if len(cmd.Arguments) == 0 {
		err := errors.New("the handlerAgg handler expects a single time argument")
		return err
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		HandlerScrapeFeeds(s)
	}
}

func HandlerScrapeFeeds(s *State) error {

	// get oldest feed
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	// mark feed fetched
	markedFeed, err := s.Db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	//fetch feed
	fetchedData, err := rss.FetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return err
	}

	//print titles to the console
	for _, v := range fetchedData.Channel.Item {
		fmt.Println(v.Title)
		fmt.Println("----------")
	}

	return nil
}
