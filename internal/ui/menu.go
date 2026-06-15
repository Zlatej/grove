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

var choices = []string{New, Delete, Update, Clone, List}

type Model struct {
	DisplayChoices []string
	Cursor         int
	Selected       string
}

func NewModel() Model {
	return Model{
		DisplayChoices: []string{"[N]ew", "[D]elete", "[U]pdate", "[C]lone", "[L]ist"},
		Cursor:         0,
		Selected:       "",
	}

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Selected = ""
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.DisplayChoices)-1 {
				m.Cursor++
			}
		case "enter", "space":
			m.Selected = choices[m.Cursor]
		case "n":
			m.Selected = New
		case "d":
			m.Selected = Delete
		case "u":
			m.Selected = Update
		case "c":
			m.Selected = Clone
		case "l":
			m.Selected = List
		}
	}
	if m.Selected != "" {
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) View() tea.View {
	var s strings.Builder
	s.WriteString("Select an action\n\n")

	for i, choice := range m.DisplayChoices {

		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		fmt.Fprintf(&s, " %s %s\n", cursor, choice)
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}
