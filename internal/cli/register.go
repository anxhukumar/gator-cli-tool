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

func HandlerRegister(s *State, cmd Command) error {
	// check if arguments are available
	if len(cmd.Arguments) == 0 {
		err := errors.New("the register handler expects a single username argument")
		return err
	}

	//user data
	userName := cmd.Arguments[0]
	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	// create user in db
	userData, err := s.Db.CreateUser(context.Background(), user)
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.ConfigPtr.SetUser(userData.Name)
	if err != nil {
		return err
	}
	fmt.Println("User created successfully")
	fmt.Println(userData)

	return nil
}
