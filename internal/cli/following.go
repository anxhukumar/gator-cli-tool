package cli

import (
	"context"
	"fmt"

	"github.com/anxhukumar/gator-cli-tool/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {

	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, v := range feeds {
		fmt.Println(v.FeedName)
		fmt.Println("---------")
	}

	return nil
}
