package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/anxhukumar/gator-cli-tool/internal/cli"
	"github.com/anxhukumar/gator-cli-tool/internal/config"
	"github.com/anxhukumar/gator-cli-tool/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// read config file
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	// sql connection
	db, err := sql.Open("postgres", conf.Db_url)

	if err != nil {
		fmt.Println("Database connection failed.", err)
		return
	}

	// create state of the current config
	confState := cli.State{
		Db:        &database.Queries{},
		ConfigPtr: &conf,
	}

	// db methods
	dbQueries := database.New(db)

	confState.Db = dbQueries

	// init commands struct
	cmdsDir := cli.Commands{
		Cmds: make(map[string]func(*cli.State, cli.Command) error),
		Mu:   &sync.RWMutex{},
	}

	// registering commands
	cmdsDir.Register("login", cli.HandlerLogin)
	cmdsDir.Register("register", cli.HandlerRegister)
	cmdsDir.Register("reset", cli.HandlerReset)
	cmdsDir.Register("users", cli.HandlerGetUsers)
	cmdsDir.Register("agg", cli.HandlerAgg)
	cmdsDir.Register("addfeed", middlewareLoggedIn(cli.HandlerAddFeed))
	cmdsDir.Register("feeds", cli.HandlerListFeeds)
	cmdsDir.Register("follow", middlewareLoggedIn(cli.HandlerFollowFeed))
	cmdsDir.Register("following", middlewareLoggedIn(cli.HandlerFollowing))
	cmdsDir.Register("unfollow", middlewareLoggedIn(cli.HandlerUnfollow))
	cmdsDir.Register("browse", middlewareLoggedIn(cli.HandlerBrowse))

	// run commands
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Error: Missing arguments | Format: login <username>")
		os.Exit(1)
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	cmd := cli.Command{
		Name:      cmdName,
		Arguments: cmdArgs,
	}

	err = cmdsDir.Run(&confState, cmd)
	if err != nil {
		fmt.Println(err)
	}

}

func middlewareLoggedIn(handler func(s *cli.State, cmd cli.Command, user database.User) error) func(*cli.State, cli.Command) error {
	return func(s *cli.State, cmd cli.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.ConfigPtr.Current_user_name)

		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
