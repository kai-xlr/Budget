package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Paycheck struct {
	Amount float64
}

type Expense struct {
	Name   string
	Amount float64
}

type Session struct {
	Date     time.Time
	Paycheck Paycheck
	Expenses []Expense
}

const dataFile = "budget_data.json"

func saveSession(s Session) error {
	sessions := loadSessions()
	sessions = append(sessions, s)
	data, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func loadSessions() []Session {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return []Session{}
	}
	var sessions []Session
	err = json.Unmarshal(data, &sessions)
	if err != nil {
		return []Session{}
	}
	return sessions
}

func payday(scanner *bufio.Scanner) {
	p := paycheck(scanner)
	expenses := expenses(scanner)

	totalExpenses := 0.0
	for _, e := range expenses {
		totalExpenses += e.Amount
	}

	fmt.Printf("Paycheck Amount: %.2f\n", p.Amount)
	fmt.Printf("Total Expenses: %.2f\n", totalExpenses)
	for _, e := range expenses {
		fmt.Printf("  - %s: %.2f\n", e.Name, e.Amount)
	}
	fmt.Printf("Remaining: %.2f\n", p.Amount-totalExpenses)

	session := Session{
		Date:     time.Now(),
		Paycheck: p,
		Expenses: expenses,
	}
	err := saveSession(session)
	if err != nil {
		fmt.Println("Error saving session:", err)
	}
}

func paycheck(scanner *bufio.Scanner) Paycheck {
	fmt.Print("Enter your paycheck amount: ")
	scanner.Scan()
	amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		fmt.Println("Invalid amount, defaulting to 0.00")
		return Paycheck{Amount: 0}
	}
	return Paycheck{Amount: amount}
}

func expenses(scanner *bufio.Scanner) []Expense {
	var expenses []Expense
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
		expenses = append(expenses, Expense{Name: name, Amount: amount})
	}
	return expenses
}

func trends() {
	sessions := loadSessions()
	if len(sessions) == 0 {
		fmt.Println("No sessions recorded yet.")
		return
	}

	fmt.Printf("\n=== Trends (%d sessions) ===\n\n", len(sessions))
	for i, s := range sessions {
		totalExpenses := 0.0
		for _, e := range s.Expenses {
			totalExpenses += e.Amount
		}
		fmt.Printf("[%d] %s\n", i+1, s.Date.Format("2006-01-02 15:04"))
		fmt.Printf("    Paycheck: %.2f, Expenses: %.2f, Remaining: %.2f\n",
			s.Paycheck.Amount, totalExpenses, s.Paycheck.Amount-totalExpenses)
	}
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Welcome to your budget app!")
		fmt.Println("1. Payday Mode")
		fmt.Println("2. Trends Mode")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			payday(scanner)
		case "2":
			trends()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}
