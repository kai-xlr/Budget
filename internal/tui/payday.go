package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kai-xlr/Budget/internal/models"
	"github.com/kai-xlr/Budget/internal/storage"
)

func updatePaycheckInput(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.paycheckInput, cmd = m.paycheckInput.Update(msg)

	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, cmd
	}

	if keyMsg.String() == "enter" {
		val := strings.TrimSpace(m.paycheckInput.Value())
		if val == "" {
			m.errMsg = "Amount cannot be empty"
			return m, cmd
		}
		amount, err := strconv.ParseFloat(val, 64)
		if err != nil {
			m.errMsg = "Invalid amount. Enter a number."
			return m, cmd
		}
		m.paycheckAmount = amount
		m.paycheckInput.Blur()
		m.expenses = nil
		m.errMsg = ""
		m.state = expenseNameInputState
		m.expenseNameInput.Focus()
		m.expenseNameInput.SetValue("")
		return m, cmd
	}

	return m, cmd
}

func updateExpenseNameInput(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.expenseNameInput, cmd = m.expenseNameInput.Update(msg)

	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, cmd
	}

	if keyMsg.String() == "enter" {
		name := strings.TrimSpace(m.expenseNameInput.Value())
		if name == "" {
			session := models.Session{
				Date:     time.Now(),
				Paycheck: models.Paycheck{Amount: m.paycheckAmount},
				Expenses: m.expenses,
			}
			storage.SaveSession(session)
			m.state = paydaySummaryState
			m.errMsg = ""
			return m, cmd
		}
		m.pendingExpenseName = name
		m.expenseNameInput.Blur()
		m.state = expenseAmountInputState
		m.expenseAmtInput.Focus()
		m.expenseAmtInput.SetValue("")
		m.errMsg = ""
		return m, cmd
	}

	return m, cmd
}

func updateExpenseAmountInput(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.expenseAmtInput, cmd = m.expenseAmtInput.Update(msg)

	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, cmd
	}

	if keyMsg.String() == "enter" {
		val := strings.TrimSpace(m.expenseAmtInput.Value())
		if val == "" {
			m.errMsg = "Amount cannot be empty"
			return m, cmd
		}
		amount, err := strconv.ParseFloat(val, 64)
		if err != nil {
			m.errMsg = "Invalid amount. Enter a number."
			return m, cmd
		}
		m.expenses = append(m.expenses, models.Expense{
			Name:   m.pendingExpenseName,
			Amount: amount,
		})
		m.expenseAmtInput.Blur()
		m.state = expenseNameInputState
		m.expenseNameInput.Focus()
		m.expenseNameInput.SetValue("")
		m.errMsg = ""
		return m, cmd
	}

	return m, cmd
}

func updatePaydaySummary(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state = menuState
	return m, nil
}

func viewPaycheckInput(m Model) string {
	var b strings.Builder
	b.WriteString(top())
	b.WriteString(center("Payday Mode"))
	b.WriteString(mid())
	b.WriteString(line(""))
	b.WriteString(line("  Enter paycheck amount:"))
	b.WriteString(line(""))
	b.WriteString(line("  " + m.paycheckInput.View()))
	if m.errMsg != "" {
		b.WriteString(line("  ! " + m.errMsg))
	}
	b.WriteString(bot())
	return b.String()
}

func viewExpenseNameInput(m Model) string {
	var b strings.Builder
	b.WriteString(top())
	b.WriteString(center("Payday Mode"))
	b.WriteString(mid())
	b.WriteString(line(""))
	b.WriteString(line(fmt.Sprintf("  Paycheck: $%.2f", m.paycheckAmount)))
	if len(m.expenses) > 0 {
		b.WriteString(mid())
		for _, e := range m.expenses {
			b.WriteString(line(fmt.Sprintf("    %-20s $%.2f", e.Name, e.Amount)))
		}
	}
	b.WriteString(mid())
	b.WriteString(line("  Enter expense name"))
	b.WriteString(line("  (or blank to finish):"))
	b.WriteString(line(""))
	b.WriteString(line("  " + m.expenseNameInput.View()))
	if m.errMsg != "" {
		b.WriteString(line("  ! " + m.errMsg))
	}
	b.WriteString(bot())
	return b.String()
}

func viewExpenseAmountInput(m Model) string {
	var b strings.Builder
	b.WriteString(top())
	b.WriteString(center("Payday Mode"))
	b.WriteString(mid())
	b.WriteString(line(""))
	b.WriteString(line(fmt.Sprintf("  Paycheck: $%.2f", m.paycheckAmount)))
	if len(m.expenses) > 0 {
		b.WriteString(mid())
		for _, e := range m.expenses {
			b.WriteString(line(fmt.Sprintf("    %-20s $%.2f", e.Name, e.Amount)))
		}
	}
	b.WriteString(mid())
	b.WriteString(line(fmt.Sprintf("  Amount for '%s':", truncate(m.pendingExpenseName, 20))))
	b.WriteString(line(""))
	b.WriteString(line("  " + m.expenseAmtInput.View()))
	if m.errMsg != "" {
		b.WriteString(line("  ! " + m.errMsg))
	}
	b.WriteString(bot())
	return b.String()
}

func viewPaydaySummary(m Model) string {
	totalExpenses := sumExpenses(m.expenses)
	remaining := m.paycheckAmount - totalExpenses

	var b strings.Builder
	b.WriteString(top())
	b.WriteString(center("Payday Summary"))
	b.WriteString(mid())
	b.WriteString(line(""))
	b.WriteString(line(fmt.Sprintf("  Paycheck:    $%.2f", m.paycheckAmount)))
	b.WriteString(line(fmt.Sprintf("  Expenses:    $%.2f", totalExpenses)))
	b.WriteString(line("  ─────────────────────────"))
	b.WriteString(line(fmt.Sprintf("  Remaining:   $%.2f", remaining)))
	if len(m.expenses) > 0 {
		b.WriteString(mid())
		for _, e := range m.expenses {
			b.WriteString(line(fmt.Sprintf("    %-20s $%.2f", e.Name, e.Amount)))
		}
	}
	b.WriteString(mid())
	b.WriteString(line("  Press any key to continue"))
	b.WriteString(bot())
	return b.String()
}
