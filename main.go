package main

import (
	"fmt"
	"log"
	"time"
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
func (p *NewPomo) PomoStart() {
	// convert seconds to minutes
	mins := p.duration * 60
	fmt.Println(mins)

	// start loop over time stated
	fmt.Printf("pomo starting for %d minutes ..\n", mins/60)
	//TODO: check to see if 1 minute so it replaces minutes with minute
	//TODO: check that it isn't a non-int value

	//TODO: find better idea for loop
	for i := mins; i >= 0; i-- {
		fmt.Printf("\rStudy Countdown: %d ", i) // \r returns to the start of line
		time.Sleep(1 * time.Second)
	}

}

func (p *NewPomo) PomoBreak(n int) {
	// convert to mins
	if n == 1 {
		n = 5
	} else if n == 2 {
		n = 10
	}
	mins := n * 60
	fmt.Println(mins)

	for i := mins; i >= 0; i-- {
		fmt.Printf("\r Break Countdown: %d ", i) // \r returns to the start of line
		time.Sleep(1 * time.Second)
	}
}

// pomoInit
// -- init starts program and loops over code until user quits
// pomoBreak
// -- starts pomo break 10min - n is time user decides

func main() {
	fmt.Println("Hello user")
	var (
		answer   string
		duration int
	)

	// loop dies if user types q
	for answer != "q" {

		fmt.Printf("\nDo you want to start a pomo: (y) yes, (n) no: ")
		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			log.Fatal(err)
		}

		// pomo loop
		if answer == "y" {
			// ask user for minutes
			fmt.Printf("time in minutes : ")
			_, err := fmt.Scanf("%d", &duration)
			if err != nil {
				log.Fatal(err)
			}
			p := &NewPomo{duration: duration}
			p.PomoStart()

			// break
			fmt.Printf("\nWould you like to start a break: (y)/(n)")
			var goBreak string
			fmt.Scanf("%s", &goBreak)

			// start break
			if goBreak == "y" {
				var num int
				fmt.Printf("How long do you want to break 5 or 10 minutes (1) 5 minutes, (2) 10 minutes : ")
				fmt.Scanf("%d", &num)

				p.PomoBreak(num)

				fmt.Println("Break finsihed!")
			} else if goBreak == "n" {
				log.Println("Have a great day :)")
				return
			}
			// break timer
		} else if answer == "n" {
			log.Println("Exiting..")
			return
		}
	}
}
