package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

const (
	New    = "new"
	Delete = "delete"
	Update = "update"
	Clone  = "clone"
	List   = "list"
)

type choice struct {
	display string
	key     string
	action  string
}

var choices = []choice{
	{"[N]ew", "n", New},
	{"[D]elete", "d", Delete},
	{"[U]pdate", "u", Update},
	{"[C]lone", "c", Clone},
	{"[L]ist", "l", List},
}

type Model struct {
	cursor   int
	Selected string
}

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
		if m.cursor < len(choices)-1 {
			m.cursor++
		}
	case "enter", "space":
		m.Selected = choices[m.cursor].action
		return m, tea.Quit
	default:
		for _, c := range choices {
			if c.key == key.String() {
				m.Selected = c.action
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	var s strings.Builder
	s.WriteString("Select an action\n\n")

	for i, choice := range choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		fmt.Fprintf(&s, " %s %s\n", cursor, choice.display)
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}
