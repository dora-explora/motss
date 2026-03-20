package main

import (
	// "fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

const (
	Bold = "\033[1m"
	Dim = "\033[2m"
	Underline = "\033[4m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Pink = "\033[35m"
	Cyan = "\033[36m"
	White = "\033[37m"
	BgGrey = "\033[100m"
	Reset = "\033[0m"
)

func (m model) Init() tea.Cmd {
	return loadDirectory
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    	case dirMsg:
    	m.threads = processDirectory(string(msg))

     	case threadMsg:
      	m.thread = string(msg)

       	case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height

    	case tea.KeyPressMsg:
        switch msg.String() {
        	case "ctrl+c", "q":
         	return m, tea.Quit

           	case "up", "k":
            if m.open {
            	if m.openscroll > 0 {
             		m.openscroll--
             	}
            } else {
           		if m.cursor > 0 {
               		m.cursor--
                 	if m.cursor == m.closedscroll - 1 {
                  		m.closedscroll--
                  	}
           		}
            }

           	case "down", "j":
            if m.open {
            	if m.openscroll < len(strings.Split(m.thread, "\n")) - m.height {
             		m.openscroll++
             	}
            } else {
           		if m.cursor < len(m.threads) - 1 {
          			m.cursor++
             		if m.cursor == m.closedscroll + m.height {
               			m.closedscroll++
               		}
           		}
        	}

           	case "space", "enter", "l":
            if m.open != true {
            	m.open = true
             	m.openscroll = 0
            	return m, loadThread(m.threads[m.cursor].id)
            }

            case "esc", "backspace", "h":
            m.open = false
        }
    }

    return m, nil
}

func (m model) View() tea.View {
    s := ""

    if !m.open {
    	for i, thread := range m.threads {
     		if i < m.closedscroll { continue }
       		if i > m.closedscroll + m.height - 1 { continue }

        	cursor := Reset + "  "
        	if m.cursor == i {
         		cursor = Reset + Pink + "> " + BgGrey
        	}

         	title := Green + thread.title
         	titlepadding := m.width / 2 - len(title)
          	if titlepadding >= 0 {
          		title += strings.Repeat(" ", titlepadding)
           	} else {
            	title = title[:m.width / 2 - 3]
             	title += "..."
            }

            date := Yellow + thread.date

            sender := Cyan + thread.sender
           	if len(sender) > m.width / 2 - 7 {
            	sender = sender[:m.width / 2 - 10] + "..."
            }
            senderpadding := m.width / 2 - len(sender) - 6
            sender += strings.Repeat(" ", senderpadding)

        	s += cursor + "| " + title + White + "| " + date + White + " | " + sender  + "\n"
         	// s += Reset + "  ┣" + strings.Repeat("━", len(title) - 4) + "╋━━━━━━━━━━╋" + strings.Repeat("━", m.width / 2 - 10) + "\n"
    	}
    } else {
    	lines := strings.Split(m.thread, "\n")
     	lastbreak := -1
     	for i, line := range lines {
         	color := Green
          	if line == "-----------------------------------------------------------------" {
             	color = White
           		lastbreak = i
           	}
      		if i < m.openscroll { continue }
        	if i == 0 {
         		s += Bold
         	} else if i - lastbreak == 2 {
          		color = Cyan
         	} else if i - lastbreak == 3 {
          		color = Yellow
          	} else if unpadded := strings.ReplaceAll(line, " ", ""); unpadded != ""  {
           		if unpadded[0] == '>' {
           			color += Dim
            	}
           	}
        	s += color + line + "\n" + Reset
      	}
    }

    v := tea.NewView(s)
    v.AltScreen = true

    return v
}
