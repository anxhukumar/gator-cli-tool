package cli

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error in resetting users: %w", err)
	}
	return nil
}
