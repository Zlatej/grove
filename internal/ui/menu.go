package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	Choices []string
	Cursor  int
	Slected string
}

func NewModel() Model {
	return Model{
		Choices: []string{"[N]ew", "[D]elete", "[U]pdate ", "[C]lone", "[L]ist"},
		Cursor:  0,
		Slected: "",
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
			m.Slected = ""
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter", "space":
			m.Slected = m.Choices[m.Cursor]
			return m, tea.Quit
		case "n":
			m.Slected = "new"
		case "d":
			m.Slected = "done"
		case "u":
			m.Slected = "update"
		case "c":
			m.Slected = "clone"
		case "l":
			m.Slected = "list"
		}

	}

	return m, nil
}

func (m Model) View() tea.View {
	var s strings.Builder
	s.WriteString("Select an action\n\n")

	for i, choice := range m.Choices {

		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		fmt.Fprintf(&s, "%s %s\n", cursor, choice)
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}
