package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	posts []string
	cursor int
}

func initialModel() model {
	return model{
		posts: []string{"post 1", "post 2", "post 3"},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyPressMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
        	if m.cursor < len(m.posts) - 1 {
        		m.cursor++
        	}
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() tea.View {
    // The header
    s := "What should we buy at the market?\n\n"

    // Iterate over our choices
    for i, post := range m.posts {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, post)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return tea.NewView(s)
}

func main() {
    program := tea.NewProgram(initialModel())
    _, err := program.Run();
    if err != nil {
        fmt.Printf("An error occured: %v", err)
        os.Exit(1)
    }
}
