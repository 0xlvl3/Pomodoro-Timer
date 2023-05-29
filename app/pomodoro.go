package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// will contain pomodoro methods

func main() {
	fmt.Println("vim-go")
}

// Items that I need to place in non-api

type PomodoroStore interface {
	StartPomodoroSession()
	StartStudySession()
	StartPomodoroBreak()
}

type Pomodoro struct {
	store *PomodoroStore
}

func NewPomodoroStore(store *PomodoroStore) *Pomodoro {
	return &Pomodoro{
		store: store,
	}
}

func Mainloop(user string) {
	var input int
	fmt.Printf("Hello %v!\n", user)
	fmt.Println("Where do you want to go")
	fmt.Println("1 - Todo")
	// todo add, todo list
	fmt.Println("2 - Pomodoro Timer")
	// pomo start, timer - for how long have preset options and an option to
	// set your own time
	fmt.Printf("Please selection a option: ")
	fmt.Scanf("%d", &input)

	if input == 1 {
		fmt.Println("pomo")
		StartStudySession(user)

		Mainloop(user)
	} else if input == 2 {
		fmt.Println("todo")

		Mainloop(user)
	}
}

// StartPomodoroBreak a new break timer
func StartPomodoroBreak(user string) {

	var duration int
	var input string
	fmt.Printf("How long do you want to break for: ")
	fmt.Scanf("%v", &duration)

	mins := duration * 60

	// start loop over time stated
	if mins <= 60 {
		fmt.Printf("pomo break starting for %d minute ..\n\n", mins/60)
	} else if mins > 60 {
		fmt.Printf("pomo break starting for %d minutes ..\n\n", mins/60)
	}

	//TODO: make quit function
	fmt.Println("q to quit at anytime")

	TimeLoop("Break", mins)

	fmt.Printf("\n\ngo to study (y) yes or (q) quit to menu: ")
	fmt.Scanf("%v", &input)

	if input == "y" {
		StartStudySession(user)
	} else if input == "q" {
		Mainloop(user)
	} else {
		// break or something
	}

	// go back to menu was here
}

// StartStudySession a new study timer
func StartStudySession(user string) {

	var duration int
	var input string

	fmt.Printf("How long do you want to study for %v: ", user)
	fmt.Scanf("%v", &duration)

	mins := duration * 60

	// start loop over time stated
	if mins <= 60 {
		fmt.Printf("pomo starting for %d minute ..\n\n", mins/60)
	} else if mins > 60 {
		fmt.Printf("pomo starting for %d minutes ..\n\n", mins/60)
	}

	//TODO: add quit function
	fmt.Println("q to quit at anytime")

	TimeLoop("Study", mins)

	fmt.Printf("\n\ngo to break (y) yes or (q) quit to menu: ")
	fmt.Scanf("%v", &input)

	if input == "y" {
		StartPomodoroBreak(user)
	} else if input == "q" {
		Mainloop(user)
	} else {
		// break or something
	}

}

// Start is our init and welcome menu
func StartPomodoroSession(user string) {

	input := " "
	for input != "q" {
		fmt.Printf("\nHello %v you want to start a pomo: \n(y) yes, \n(q) quit to menu: ", user)
		fmt.Scanf("%s", &input)

		if input == "y" {
			StartStudySession(user)

		} else if input == "q" {
			Mainloop(user)

		} else {
			break
		}
	}
}

func ReadUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// main loop handle case of looping
func TimeLoop(label string, minutes int) {
	fmt.Printf("%s starting for %d minutes .. ", label, minutes)

	for i := minutes; i >= 0; i-- {
		fmt.Printf("\r%s Break Countdown...: %d ", label, i) // \r returns to the start of line
		time.Sleep(1 * time.Second)
	}

}
