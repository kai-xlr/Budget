package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func updateTrends(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state = menuState
	return m, nil
}

func viewTrends(m Model) string {
	var b strings.Builder
	b.WriteString(top())
	b.WriteString(center(fmt.Sprintf("Trends (%d sessions)", len(m.sessions))))
	b.WriteString(mid())

	if len(m.sessions) == 0 {
		b.WriteString(line(""))
		b.WriteString(line("  No sessions recorded yet."))
	} else {
		for i, s := range m.sessions {
			totalExpenses := sumExpenses(s.Expenses)
			remaining := s.Paycheck.Amount - totalExpenses
			b.WriteString(line(""))
			b.WriteString(line(fmt.Sprintf("  [%d] %s", i+1, s.Date.Format("2006-01-02 15:04"))))
			b.WriteString(line(fmt.Sprintf("      Paycheck: $%.2f", s.Paycheck.Amount)))
			b.WriteString(line(fmt.Sprintf("      Expenses: $%.2f", totalExpenses)))
			b.WriteString(line(fmt.Sprintf("      Remaining: $%.2f", remaining)))
		}
	}

	b.WriteString(mid())
	b.WriteString(line("  Press any key to go back"))
	b.WriteString(bot())
	return b.String()
}
