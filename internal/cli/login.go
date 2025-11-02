package cli

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	// check if arguments are available
	if len(cmd.Arguments) == 0 {
		err := errors.New("the login handler expects a single username argument")
		return err
	}

	// set username in the state config
	userName := cmd.Arguments[0]
	err := s.ConfigPtr.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Println("New user is set")

	return nil
}
