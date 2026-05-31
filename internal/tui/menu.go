package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kai-xlr/Budget/internal/storage"
)

func updateMenu(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch keyMsg.String() {
	case "1":
		m.state = paycheckInputState
		m.paycheckInput.Focus()
		m.paycheckInput.SetValue("")
		m.errMsg = ""
		return m, nil
	case "2":
		m.sessions = storage.LoadSessions()
		m.state = trendsState
		return m, nil
	case "3", "q":
		return m, tea.Quit
	}

	return m, nil
}

func viewMenu(m Model) string {
	var b strings.Builder
	b.WriteString(top())
	b.WriteString(center("Welcome to Budget App!"))
	b.WriteString(mid())
	b.WriteString(line(""))
	b.WriteString(line("  1. Payday Mode"))
	b.WriteString(line("  2. Trends Mode"))
	b.WriteString(line("  3. Exit"))
	b.WriteString(line(""))
	b.WriteString(line("  Press 1-3 to choose"))
	b.WriteString(bot())
	return b.String()
}
