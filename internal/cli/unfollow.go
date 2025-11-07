package cli

import (
	"context"
	"errors"

	"github.com/anxhukumar/gator-cli-tool/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	// check if arguments are available
	if len(cmd.Arguments) == 0 {
		err := errors.New("the unfollow command expects a single url argument")
		return err
	}

	//get feed if from url
	feed, err := s.Db.GetFeedByUrl(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	//delete data
	err = s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	return nil
}
