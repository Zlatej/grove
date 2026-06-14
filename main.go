package main

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/zlatej/grove/internal/config"
	"github.com/zlatej/grove/internal/ui"
)

func main() {
	_, err := config.Load()
	if err != nil {
		slog.Error("loading config", "error", err.Error())
	}
	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		slog.Error("running menu model", "error", err)
	}
}
