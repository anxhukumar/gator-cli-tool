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

func HandlerAddFeed(s *State, cmd Command) error {
	// check if arguments are available
	if len(cmd.Arguments) < 2 {
		os.Exit(1)
		err := errors.New("the addFeed handler expects a name and url argument")
		return err
	}

	feedNameArg := cmd.Arguments[0]
	urlArg := cmd.Arguments[1]

	currentUser, err := s.Db.GetUser(context.Background(), s.ConfigPtr.Current_user_name)
	if err != nil {
		return err
	}

	feedData := database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedNameArg,
		Url:       urlArg,
		UserID:    currentUser.ID,
	}

	feed, err := s.Db.CreateFeeds(context.Background(), feedData)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
