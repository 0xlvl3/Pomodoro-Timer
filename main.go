package main

// ideas
// loop handled by own function

// add todo to aspect where user can add a list of todo
// todo then can be studied through pomo

// mongodb - user accs, saves data
// don't require mongodb can be optional

// understand a bit more on styling through lipgloss
// - figure way that it won't reprint on the same line with \r "lipgloss"

// bring in tea model
// has selections for pre-defined study times 15, 25, 40, 60
// has slected for pre-defined break times 5, 10, 15
// todos have own selection board

// clean up code

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	input    string
	duration int
)

type Pomo interface {
	PomoStart()
	PomoBreak(int)
}

type NewPomo struct {
	duration int
}

// pomoStart
// -- starts a fresh pomo - n is time user desires
func (p *NewPomo) PomoBegin() {

	// convert seconds to minutes
	mins := p.duration * 60
	fmt.Println(mins)

	// start loop over time stated
	fmt.Printf("pomo starting for %d minutes ..\n", mins/60)
	//TODO: check to see if 1 minute so it replaces minutes with minute
	//TODO: check that it isn't a non-int value

	//TODO: find better idea for loop
	for i := mins; i >= 0; i-- {
		fmt.Printf("\rStudy Countdown: %d", i) // \r returns to the start of line
		time.Sleep(1 * time.Second)
	}

}

func (p *NewPomo) Start() {

	fmt.Println("Hello user")

	for input != "q" {
		fmt.Printf("\nDo you want to start a pomo: (y) yes, (t) todo, (q) quit: ")

		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		switch input {
		case "y":
			p.PomoBegin()
		// start
		case "t":
			p.Todo()
		// create a todo
		case "q":
			// user quits applicaton
			os.Exit(2)
		}
		fmt.Println("\nwhere now?")
	}
}

func (p *NewPomo) Todo() {
	fmt.Println("todo")

}

// main loop handle case of looping
func (p *NewPomo) TimeLoop(userInput string, t int) {

	switch userInput {

	// if user starts pomo
	case "y":
		mins := t * 60
		fmt.Println(mins)

		for i := mins; i >= 0; i-- {
			fmt.Printf("\r Study Countdown...: %d ", i) // \r returns to the start of line
			time.Sleep(1 * time.Second)
		}
	case "b":
		mins := t * 60
		fmt.Println(mins)

		for i := mins; i >= 0; i-- {
			fmt.Printf("\r Break Countdown...: %d ", i) // \r returns to the start of line
			time.Sleep(1 * time.Second)
		}

	}
}

func main() {

	fmt.Println("lol")
	p := NewPomo{}
	p.Start()

}
