package cli

import (
	"context"
	"fmt"

	"github.com/anxhukumar/gator-cli-tool/internal/rss"
)

const url string = "https://www.wagslane.dev/index.xml"

func HandlerAgg(s *State, cmd Command) error {
	rssFeed, err := rss.FetchFeed(context.Background(), url)

	if err != nil {
		return err
	}

	fmt.Println(rssFeed)

	return nil
}
