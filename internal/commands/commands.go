package commands

import (
	"fmt"

	"github.com/zlatej/grove/internal/config"
	"github.com/zlatej/grove/internal/git"
)

// Choice contains information about one command
type Choice struct {
	Display string
	Key     string
	ID      CommandID
}

// Choices contain all implemented commands user can pick from
var Choices = []Choice{
	{"[N]ew", "n", New},
	{"[D]elete", "d", Delete},
	{"[U]pdate", "u", Update},
	{"[C]lone", "c", Clone},
	{"[L]ist", "l", List},
}

type CommandID string

const (
	New    CommandID = "new"
	Delete CommandID = "delete"
	Update CommandID = "update"
	Clone  CommandID = "clone"
	List   CommandID = "list"
)

type commandFunc func(cfg *config.Config) error

var registry = map[CommandID]commandFunc{
	New:    runNew,
	Delete: runDelete,
	Update: runUpdate,
	Clone:  runClone,
	List:   runList,
}

func RunCommand(cmdID CommandID, cfg *config.Config) error {
	cmd, ok := registry[cmdID]
	if !ok {
		return fmt.Errorf("command with identifier %s does not exist", cmdID)
	}
	return cmd(cfg)
}

func runNew(cfg *config.Config) error {
	return nil
}

func runDelete(cfg *config.Config) error {
	return nil
}

func runUpdate(cfg *config.Config) error {
	return nil
}

func runClone(cfg *config.Config) error {
	url, err := RepoUrl()
	if err != nil {
		return err
	}
	err = git.CloneBare(url, "")
	if err != nil {
		return err
	}
	return nil
}

func runList(cfg *config.Config) error {
	return nil
}
