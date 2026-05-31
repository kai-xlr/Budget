package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kai-xlr/Budget/internal/budget"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Welcome to your budget app!")
		fmt.Println("1. Payday Mode")
		fmt.Println("2. Trends Mode")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			budget.Payday(scanner)
		case "2":
			budget.Trends()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}
