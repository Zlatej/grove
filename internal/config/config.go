package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const GroveConfig = "grove/config.toml"

type Config struct {
	TicketIDPrefix string   `toml:"ticket_id_prefix"`
	TicketIDFirst  bool     `toml:"ticket_id_first"`
	BaseBranch     string   `toml:"base_branch"`
	CloneWT        []string `toml:"clone_wt"`
}

func Load() (*Config, error) {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("getting config path: %w", err)
	}

	path := filepath.Join(cfgPath, GroveConfig)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = createNew(path); err != nil {
			return nil, fmt.Errorf("creating new config: %w", err)
		}
	}

	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("decoding config: %w", err)
	}

	return &cfg, nil
}

func createNew(path string) error {
	newCfg := []byte(`# this is default config, adjust it to your needs

# if you require to have the id of a ticket in branch name, uncomment and enter your ID prefix
# ticket_id_prefix = "TID-" 
ticket_id_first = true

base_branch = "main"

# worktrees created after a repository is cloned
clone_wt = ["main", "review"]`)

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err = f.Write(newCfg); err != nil {
		return err
	}
	return f.Close()
}
