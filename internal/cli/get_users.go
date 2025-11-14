package cli

import (
	"context"
	"fmt"
)

func HandlerGetUsers(s *State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, v := range users {
		if v.Name == s.ConfigPtr.Current_user_name {
			fmt.Printf("* %v (current)\n", v.Name)
		} else {
			fmt.Printf("* %v\n", v.Name)
		}
	}

	return nil
}
