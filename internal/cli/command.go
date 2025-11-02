package cli

import (
	"errors"
	"sync"
)

type Command struct {
	Name      string
	Arguments []string
}

// collection of all Commands in map with its function call
type Commands struct {
	Cmds map[string]func(*State, Command) error
	Mu   *sync.RWMutex
}

func (c *Commands) Run(s *State, cmd Command) error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	v, ok := c.Cmds[cmd.Name]
	if !ok {
		err := errors.New("Command does not exist")
		return err
	}
	v(s, cmd)

	return nil
}

func (c *Commands) Register(Name string, f func(*State, Command) error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Cmds[Name] = f
}
