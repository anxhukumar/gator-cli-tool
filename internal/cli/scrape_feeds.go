package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/anxhukumar/gator-cli-tool/internal/database"
	"github.com/anxhukumar/gator-cli-tool/internal/rss"
	"github.com/google/uuid"
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

	for _, v := range fetchedData.Channel.Item {

		t, _ := time.Parse(time.RFC1123Z, v.PubDate)

		data := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: sql.NullString{
				String: v.Title,
				Valid:  true,
			},
			Url: v.Link,
			Description: sql.NullString{
				String: v.Description,
				Valid:  true,
			},
			PublishedAt: t,
			FeedID:      feed.ID,
		}

		// create posts
		_, err := s.Db.CreatePost(context.Background(), data)
		if err != nil {
			return err
		}
	}

	return nil
}
