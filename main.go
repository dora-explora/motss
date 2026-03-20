package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
    program := tea.NewProgram(initialModel())
    _, err := program.Run();
    if err != nil {
        fmt.Printf("An error occured: %v", err)
        os.Exit(1)
    }
}
