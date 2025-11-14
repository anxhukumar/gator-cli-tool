package cli

import (
	"context"
	"fmt"
)

func HandlerListFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, v := range feeds {
		userData, err := s.Db.GetUserFromId(context.Background(), v.UserID)
		if err != nil {
			return err
		}
		fmt.Println(v.Name)
		fmt.Println(v.Url)
		fmt.Println(userData.Name)
		fmt.Println("------------")
	}

	return nil
}
