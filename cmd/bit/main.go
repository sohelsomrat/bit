package main

import (
	"fmt"
	"os"

	"github.com/superstarryeyes/bit/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m, err := ui.InitialModel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
