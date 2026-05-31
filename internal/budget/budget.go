package budget

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kai-xlr/Budget/internal/models"
	"github.com/kai-xlr/Budget/internal/storage"
)

func Payday(scanner *bufio.Scanner) {
	p := paycheck(scanner)
	expenses := expenses(scanner)

	totalExpenses := sum(expenses)

	fmt.Printf("Paycheck Amount: %.2f\n", p.Amount)
	fmt.Printf("Total Expenses: %.2f\n", totalExpenses)
	for _, e := range expenses {
		fmt.Printf("  - %s: %.2f\n", e.Name, e.Amount)
	}
	fmt.Printf("Remaining: %.2f\n", p.Amount-totalExpenses)

	session := models.Session{
		Date:     time.Now(),
		Paycheck: p,
		Expenses: expenses,
	}
	err := storage.SaveSession(session)
	if err != nil {
		fmt.Println("Error saving session:", err)
	}
}

func Trends() {
	sessions := storage.LoadSessions()
	if len(sessions) == 0 {
		fmt.Println("No sessions recorded yet.")
		return
	}

	fmt.Printf("\n=== Trends (%d sessions) ===\n\n", len(sessions))
	for i, s := range sessions {
		totalExpenses := sum(s.Expenses)
		fmt.Printf("[%d] %s\n", i+1, s.Date.Format("2006-01-02 15:04"))
		fmt.Printf("    Paycheck: %.2f, Expenses: %.2f, Remaining: %.2f\n",
			s.Paycheck.Amount, totalExpenses, s.Paycheck.Amount-totalExpenses)
	}
}

func paycheck(scanner *bufio.Scanner) models.Paycheck {
	fmt.Print("Enter your paycheck amount: ")
	scanner.Scan()
	amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		fmt.Println("Invalid amount, defaulting to 0.00")
		return models.Paycheck{Amount: 0}
	}
	return models.Paycheck{Amount: amount}
}

func expenses(scanner *bufio.Scanner) []models.Expense {
	var expenses []models.Expense
	for {
		fmt.Print("Enter expense name (or blank to finish): ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())
		if name == "" {
			break
		}
		fmt.Print("Enter expense amount: ")
		scanner.Scan()
		amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("Invalid amount, skipping this expense.")
			continue
		}
		expenses = append(expenses, models.Expense{Name: name, Amount: amount})
	}
	return expenses
}

func sum(expenses []models.Expense) float64 {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	return total
}
