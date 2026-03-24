package main

import (
	// "fmt"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
)

const (
	Bold = "\033[1m"
	Dim = "\033[2m"
	Underline = "\033[4m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Pink = "\033[35m"
	Cyan = "\033[36m"
	White = "\033[37m"
	BgGrey = "\033[100m"
	Reset = "\033[0m"
)

func (m model) Init() tea.Cmd {
	return loadInit
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    	case initMsg:
    	m.threads = processDirectory(msg[0])
     	m.intro = msg[1]

     	case threadMsg:
      	m.thread = string(msg)

       	case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height

    	case tea.KeyPressMsg:
     	if m.intro != "" && m.width >= 80 && m.height >= 25 {
      		m.intro = ""
      	} else {
       		switch msg.String() {
        		case "q", "esc", "ctrl+c":
         		return m, tea.Quit

           		case "up", "k", "w":
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

           		case "down", "j", "s":
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

           		case "right", "l", "d":
            	if m.open != true {
            		m.open = true
             		m.openscroll = 0
            		return m, loadThread(m.threads[m.cursor].id)
            	}

            	case "left", "h", "a":
            	m.open = false
         	}
        }
    }

    return m, nil
}

func (m model) View() tea.View {
    s := ""

    if m.width < 80 || m.height < 25 {
    	s = m.undersizedView()
    } else if m.intro != "" {
    	s = m.introView()
    } else if m.open {
    	s = m.openView()
    } else {
    	s = m.closedView()
    }

    v := tea.NewView(s)
    v.AltScreen = true

    return v
}

func (m model) introView() string {
	str := strings.Repeat("\n", max(m.height - 25, 0) / 2)
    lines := strings.Split(m.intro, "\n")
    padding := strings.Repeat(" ", max(m.width - 70, 0) / 2)
    for i, line := range lines {
    color := Reset + Green
        if i == 9 {
        	color = Yellow + Bold
        } else if i > 18 {
        	color = Cyan
        }
        str += color + padding + line + "\n"
    }
    return str
}

func (m model) closedView() string {
	str := ""
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

        str += cursor + "| " + title + White + "| " + date + White + " | " + sender  + "\n"
        // s += Reset + "  ┣" + strings.Repeat("━", len(title) - 4) + "╋━━━━━━━━━━╋" + strings.Repeat("━", m.width / 2 - 10) + "\n"
    }
    return str
}

func (m model) openView() string {
	str := ""
	lines := strings.Split(m.thread, "\n")
    lastbreak := -1
    padding := strings.Repeat(" ", (m.width - 80) / 2)
    for i, line := range lines {
        color := Green
        if line == "-----------------------------------------------------------------" {
            color = White
           	lastbreak = i
        }
        if i < m.openscroll { continue }
        if i == 0 {
        	str += Bold
        } else if i - lastbreak == 2 {
        	color = Cyan
        } else if i - lastbreak == 3 {
        	color = Yellow
        } else if unpadded := strings.ReplaceAll(line, " ", ""); unpadded != ""  {
          		if unpadded[0] == '>' {
          			color += Dim
           	}
        }
        str += color + padding + line + "\n" + Reset
    }
    return str
}

func (m model) undersizedView() string {
	hpadding := (m.width - 27) / 2
	vpadding := (m.height - 6) / 2
	str := strings.Repeat("\n", vpadding)
	str += strings.Repeat(" ", hpadding)
	str += "Your terminal is too small!\n"
	str += strings.Repeat(" ", hpadding)
	str += "   Your terminal's size:   \n"
	str += strings.Repeat(" ", hpadding)
	str += "          "
	if m.width < 80 { str += Red } else { str += Green }
	str += strconv.Itoa(m.width) + Reset + " x "
	if m.height < 25 { str += Red } else { str += Green }
	str += strconv.Itoa(m.height) + "          \n\n" + Reset
	str += strings.Repeat(" ", hpadding)
	str += "  This app needs 80 x 25.  \n"
	str += strings.Repeat(" ", hpadding)
	str += "Exit with q, Ctrl+C, or Esc\n"
	str += strings.Repeat("\n", vpadding)
	return str
}
