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
	closedscroll int
	open bool
	openscroll int
	thread string
	width int
	height int
}

func initialModel() model {
	return model{
		threads: []thread{},
		cursor: 0,
		closedscroll: 0,
		open: false,
		openscroll: 0,
		thread: "",
		width: 150,
		height: 50,
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
