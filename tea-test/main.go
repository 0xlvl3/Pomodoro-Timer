package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const url = "https://charm.sh"

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	PaddingTop(2).
	PaddingLeft(4).
	Width(22)

// model ->
type model struct {
	status int
	err    error
}

// passed into our update
type statusMsg int
type errMsg struct{ err error }

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

// cmd and msgs
// cmds are functions that perform some I/O and then return a msg
// functions such as :
// -checking time
// -ticker a timer
// -reading from the disk
// -network stuff
// are all I/O and should be run through commands.
// this will keep it straight forward and simple
func checkServer() tea.Msg {

	// create a http client and make a GET resp
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}

	return statusMsg(res.StatusCode)

}

// return our Cmd here
func (m model) Init() tea.Cmd {

	return checkServer
}

// update method Cmds run asynchronously in a go routine
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case statusMsg:
		// update with our model with the status message
		m.status = int(msg)
		return m, tea.Quit
	case errMsg:
		// if there was an error in the model
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		// KeyMsg represents our quit
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	}
	// if we happen to get another message don't do anything
	return m, nil
}

// look at current model and build a string accordingly

func (m model) View() string {

	// if err print and don't do anything else
	if m.err != nil {
		return fmt.Sprintf("\nWe had trouble: %v\n\n", m.err)
	}

	s := fmt.Sprintf("checking %s ...", url)

	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}

	return "\n" + s + "\n\n"
}
func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
