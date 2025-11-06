package cli

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.ConfigPtr.Current_user_name)
	if err != nil {
		return err
	}

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
