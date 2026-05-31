package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/kai-xlr/Budget/internal/models"
)

type state int

const (
	menuState state = iota
	paycheckInputState
	expenseNameInputState
	expenseAmountInputState
	paydaySummaryState
	trendsState
)

const boxWidth = 44

func top() string {
	return "╔" + strings.Repeat("═", boxWidth-2) + "╗\n"
}

func mid() string {
	return "╠" + strings.Repeat("═", boxWidth-2) + "╣\n"
}

func bot() string {
	return "╚" + strings.Repeat("═", boxWidth-2) + "╝\n"
}

func line(s string) string {
	pad := boxWidth - 4 - len(s)
	if pad < 0 {
		pad = 0
	}
	return "║  " + s + strings.Repeat(" ", pad) + " ║\n"
}

func center(s string) string {
	pad := boxWidth - 4 - len(s)
	if pad < 0 {
		pad = 0
	}
	l := pad / 2
	r := pad - l
	return "║" + strings.Repeat(" ", l) + s + strings.Repeat(" ", r) + "║\n"
}

func truncate(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}

func sumExpenses(expenses []models.Expense) float64 {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	return total
}

type Model struct {
	state              state
	paycheckInput      textinput.Model
	expenseNameInput   textinput.Model
	expenseAmtInput    textinput.Model
	paycheckAmount     float64
	expenses           []models.Expense
	sessions           []models.Session
	errMsg             string
	pendingExpenseName string
}

func New() Model {
	pi := textinput.New()
	pi.Placeholder = "e.g. 5000.00"
	pi.Prompt = "> "
	pi.CharLimit = 20

	eni := textinput.New()
	eni.Placeholder = "e.g. Rent"
	eni.Prompt = "> "
	eni.CharLimit = 50

	eai := textinput.New()
	eai.Placeholder = "e.g. 1500.00"
	eai.Prompt = "> "
	eai.CharLimit = 20

	return Model{
		state:            menuState,
		paycheckInput:    pi,
		expenseNameInput: eni,
		expenseAmtInput:  eai,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.state != menuState {
				m.state = menuState
				return m, nil
			}
		}
	}

	switch m.state {
	case menuState:
		return updateMenu(m, msg)
	case paycheckInputState:
		return updatePaycheckInput(m, msg)
	case expenseNameInputState:
		return updateExpenseNameInput(m, msg)
	case expenseAmountInputState:
		return updateExpenseAmountInput(m, msg)
	case paydaySummaryState:
		return updatePaydaySummary(m, msg)
	case trendsState:
		return updateTrends(m, msg)
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case menuState:
		return viewMenu(m)
	case paycheckInputState:
		return viewPaycheckInput(m)
	case expenseNameInputState:
		return viewExpenseNameInput(m)
	case expenseAmountInputState:
		return viewExpenseAmountInput(m)
	case paydaySummaryState:
		return viewPaydaySummary(m)
	case trendsState:
		return viewTrends(m)
	}
	return ""
}
