package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/anxhukumar/gator-cli-tool/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2

	if len(cmd.Arguments) == 1 {
		limit, _ = strconv.Atoi(cmd.Arguments[0])
	}

	// get posts from userId and limit
	postData, err := s.Db.GetPostsFromUser(context.Background(), database.GetPostsFromUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("error in getting posts: %w", err)
	}

	// print the posts
	for _, v := range postData {
		fmt.Println("---------------")
		fmt.Println("Title:", v.Title.String)
		fmt.Println("Feed name:", v.FeedName)
		fmt.Println("Url:", v.Url)
		fmt.Println("Description:", v.Description.String)
	}

	return nil

}
