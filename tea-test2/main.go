package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var header = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#7D56F4")).
	Padding(2).
	Width(50).Align(lipgloss.Center)

var itemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#5AEAE9"))

// Init, function that returns initial command for the application to run
// Update, a function that handles incoming events and updates the model
// View, a function that renders UI based on the data in the model

// model stores application state
type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

// init our model with values
func initialModel() model {
	return model{
		// list of items.
		choices: []string{"buy carrots", "buy celery", "buy kohlrabi"},

		// used to select choices.
		selected: make(map[int]struct{}),
	}
}

// perform some init I/O
func (m model) Init() tea.Cmd {
	// no I/O right now
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// is it a key press?
	case tea.KeyMsg:

		// which key was pressed
		switch msg.String() {

		// quit program
		case "crtl+c", "q":
			return m, tea.Quit

			// "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

			// "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// toggle with enter and space on selected item
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}

	}

	// return updated model to process
	return m, nil

}

// where we render our UI
func (m model) View() string {
	// the header
	s := "What should we buy at the market?\n\n"

	// iterate over our choices
	for i, choice := range m.choices {

		// is the cursor pointing at this choice
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}

		// is this choice selected ?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected
		}

		//render the row
		s += itemStyle.Render(fmt.Sprintf("\n%s [%s] %s\n", cursor, checked, choice))

	}

	// the footer
	s += "\n\nPress q to quit.\n"

	// send the UI for rendering
	return s
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
