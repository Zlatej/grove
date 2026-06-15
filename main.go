package main

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/zlatej/grove/internal/commands"
	"github.com/zlatej/grove/internal/config"
	"github.com/zlatej/grove/internal/ui"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("loading config", "error", err)
		return
	}
	p := tea.NewProgram(ui.NewModel())
	tm, err := p.Run()
	if err != nil {
		slog.Error("running menu model", "error", err)
		return
	}
	m, ok := tm.(ui.Model)
	if !ok || len(m.Selected) == 0 {
		return
	}
	if err := commands.RunCommand(m.Selected, cfg); err != nil {
		slog.Error("executing command", "command", m.Selected, "error", err)
	}
}
