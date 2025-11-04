package cli

import (
	"github.com/anxhukumar/gator-cli-tool/internal/config"
	"github.com/anxhukumar/gator-cli-tool/internal/database"
)

type State struct {
	Db        *database.Queries
	ConfigPtr *config.Config
}
