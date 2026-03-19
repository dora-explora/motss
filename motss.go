package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"

	tea "charm.land/bubbletea/v2"
)

type thread struct {
	id int
	title string
	sender string
	date string
}

type model struct {
	threads []thread
	cursor int
	open bool
	scroll int
	thread string
}


func initialModel() model {
	return model{
		threads: []thread{},
		cursor: 0,
		open: false,
		scroll: 0,
		thread: "",
	}
}

func loadDirectory() tea.Msg {
	path := "./archive/directory.txt"
	contents, error := os.ReadFile(path)
	if error != nil {
		panic(error)
	}
	return dirMsg(string(contents))
}

func processDirectory(str string) []thread {
	var threads []thread
	lines := strings.Split(str, "\n")
	for i := 0; i < len(lines) - 1; i += 5 {
		id, err := strconv.Atoi(lines[i])
		if err != nil { panic(err) }
		title := lines[i + 1]
		sender := lines[i + 2]
		date := lines[i + 3]
		thread := thread {
			id: id,
			title: title,
			sender: sender,
			date: date,
		}
		threads = append(threads, thread)
	}
	return threads
}

type dirMsg string

func loadThread(id int) tea.Cmd {
	return func() tea.Msg {
		path := fmt.Sprintf("./archive/threads/%d.txt", id)
		contents, error := os.ReadFile(path)
		if error != nil {
			panic(error)
		}
		return threadMsg(string(contents))
	}
}

type threadMsg string

func (m model) Init() tea.Cmd {
	return loadDirectory
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    	case dirMsg:
    	m.threads = processDirectory(string(msg))

     	case threadMsg:
      	m.thread = string(msg)

    	case tea.KeyPressMsg:
        switch msg.String() {
        	case "ctrl+c", "q":
         	return m, tea.Quit

           	case "up", "k":
           	if m.cursor > 0 {
               	m.cursor--
           	}

           	case "down", "j":
           	if m.cursor < len(m.threads) - 1 {
           		m.cursor++
           	}

           	case "space", "enter":
            m.open = true
            return m, loadThread(m.threads[m.cursor].id)

            case "shift":
            m.thread = "AAAAAAAAA"
            return m, nil
        }
    }

    return m, nil
}

func (m model) View() tea.View {
    s := ""

    if !m.open {
    	for i, thread := range m.threads {

        	cursor := " "
        	if m.cursor == i {
            	cursor = ">"
        	}

        	s += fmt.Sprintf("%s %s - %s\n", cursor, thread.title, thread.date)
    	}
    } else {
    	s = m.thread
    }

    s += "\nPress q to quit.\n"

    v := tea.NewView(s)
    v.AltScreen = true

    return v
}

func main() {
    program := tea.NewProgram(initialModel())
    _, err := program.Run();
    if err != nil {
        fmt.Printf("An error occured: %v", err)
        os.Exit(1)
    }
}
