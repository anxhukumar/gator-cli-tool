package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/anxhukumar/gator-cli-tool/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	// check if arguments are available
	if len(cmd.Arguments) < 2 {
		os.Exit(1)
		err := errors.New("the addFeed handler expects a name and url argument")
		return err
	}

	feedNameArg := cmd.Arguments[0]
	urlArg := cmd.Arguments[1]

	feedData := database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedNameArg,
		Url:       urlArg,
		UserID:    user.ID,
	}

	feed, err := s.Db.CreateFeeds(context.Background(), feedData)
	if err != nil {
		return err
	}

	// creating a feed follow record
	feed_follow_data := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), feed_follow_data)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
