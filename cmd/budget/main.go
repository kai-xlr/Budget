package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kai-xlr/Budget/internal/budget"
	"github.com/kai-xlr/Budget/internal/gui"
	"github.com/kai-xlr/Budget/internal/tui"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║      Budget App - Launcher       ║")
		fmt.Println("╠══════════════════════════════════╣")
		fmt.Println("║                                  ║")
		fmt.Println("║  1. Simple CLI                   ║")
		fmt.Println("║  2. Terminal UI (TUI)            ║")
		fmt.Println("║  3. GUI                          ║")
		fmt.Println("║  4. Exit                         ║")
		fmt.Println("║                                  ║")
		fmt.Println("╚══════════════════════════════════╝")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			runCLI(scanner)
		case "2":
			runTUI()
		case "3":
			gui.Launch()
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}
}

func runCLI(scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║        Budget App - CLI          ║")
		fmt.Println("╠══════════════════════════════════╣")
		fmt.Println("║                                  ║")
		fmt.Println("║  1. Payday Mode                  ║")
		fmt.Println("║  2. Trends Mode                  ║")
		fmt.Println("║  3. Back to Launcher             ║")
		fmt.Println("║                                  ║")
		fmt.Println("╚══════════════════════════════════╝")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			budget.Payday(scanner)
		case "2":
			budget.Trends()
		case "3":
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}

func runTUI() {
	p := tea.NewProgram(tui.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
	}
}
