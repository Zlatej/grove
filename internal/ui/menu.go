package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/zlatej/grove/internal/commands"
)

type Model struct {
	cursor   int
	Selected commands.CommandID
}

// NewModel returns default Model, kept in case non primitive fields were added
func NewModel() Model {
	return Model{
		cursor:   0,
		Selected: "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}
	switch key.String() {
	case "ctrl+c", "q":
		m.Selected = ""
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(commands.Choices)-1 {
			m.cursor++
		}
	case "enter", "space":
		m.Selected = commands.Choices[m.cursor].ID
		return m, tea.Quit
	default:
		for _, c := range commands.Choices {
			if c.Key == key.String() {
				m.Selected = c.ID
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	var s strings.Builder
	s.WriteString("Select an action\n\n")

	for i, c := range commands.Choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		fmt.Fprintf(&s, " %s %s\n", cursor, c.Display)
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}
