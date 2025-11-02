package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/anxhukumar/gator-cli-tool/internal/cli"
	"github.com/anxhukumar/gator-cli-tool/internal/config"
)

func main() {
	// read config file
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	// create state of the current config
	confState := cli.State{
		ConfigPtr: &conf,
	}

	// init commands struct
	cmdsDir := cli.Commands{
		Cmds: make(map[string]func(*cli.State, cli.Command) error),
		Mu:   &sync.RWMutex{},
	}

	cmdsDir.Register("login", cli.HandlerLogin)

	// get arguments
	cliArgs := os.Args[1:]

	if len(cliArgs) < 2 {
		fmt.Println("Error: Missing arguments | Format: login <username>")
		os.Exit(1)
	}

	// run commands
	cmd := cli.Command{
		Name:      cliArgs[0],
		Arguments: []string{cliArgs[1]},
	}

	err = cmdsDir.Run(&confState, cmd)
	if err != nil {
		fmt.Println(err)
	}

}
