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

func HandlerFollowFeed(s *State, cmd Command) error {
	// check if arguments are available
	if len(cmd.Arguments) == 0 {
		os.Exit(1)
		err := errors.New("the follow command expects a single url argument")
		return err
	}

	user, err := s.Db.GetUser(context.Background(), s.ConfigPtr.Current_user_name)
	if err != nil {
		return err
	}

	feed, err := s.Db.GetFeedByUrl(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	feed_follow_data := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollowRes, err := s.Db.CreateFeedFollow(context.Background(), feed_follow_data)
	if err != nil {
		return err
	}

	fmt.Println(feedFollowRes.FeedName)
	fmt.Println(feedFollowRes.UserName)

	return nil
}
