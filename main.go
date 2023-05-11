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
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0xlvl3/pomodoro-timer/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongo

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

func (p *NewPomo) PomoBreak() {
	fmt.Printf("How long do you want to break for: ")
	fmt.Scanf("%v", &p.duration)

	// convert seconds to minutes
	mins := p.duration * 60

	// start loop over time stated
	fmt.Printf("pomo break for %d minutes ..\n\n", mins/60)
	fmt.Println("q to quit at anytime")
	//TODO: check to see if 1 minute so it replaces minutes with minute
	//TODO: check that it isn't a non-int value

	//TODO: find better idea for loop
	for i := mins; i >= 0; i-- {
		fmt.Printf("\rBreak Countdown: %d", i) // \r returns to the start of line
		time.Sleep(1 * time.Second)

	}

	fmt.Printf("\n\ngo to study (y) yes, (m) menu or (q) quit: ")
	fmt.Scanf("%v", &input)
	switch input {
	case "y":
		// study
		p.PomoStudy()
	case "m":
		//menu
		p.Start()
	case "q":
		//quit
		fmt.Println("quitting...")
		os.Exit(4)
	}

}

// pomoStart
// -- starts a fresh pomo - n is time user desires
func (p *NewPomo) PomoStudy() {
	fmt.Printf("How long do you want to study for: ")
	fmt.Scanf("%v", &p.duration)

	// convert seconds to minutes
	mins := p.duration * 60

	// start loop over time stated
	fmt.Printf("pomo starting for %d minutes ..\n\n", mins/60)
	fmt.Println("q to quit at anytime")
	//TODO: check to see if 1 minute so it replaces minutes with minute
	//TODO: check that it isn't a non-int value

	//TODO: find better idea for loop
	for i := mins; i >= 0; i-- {
		fmt.Printf("\rStudy Countdown: %d", i) // \r returns to the start of line
		time.Sleep(1 * time.Second)

	}

	fmt.Printf("\n\ngo to break (y) yes, (m) menu or (q) quit: ")
	fmt.Scanf("%v", &input)
	switch input {
	case "y":
		//break
		p.PomoBreak()
	case "m":
		//menu
		p.Start()
	case "q":
		//quit
		fmt.Println("quitting...")
		os.Exit(3)
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
			p.PomoStudy()
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(db.DBNAME).Collection("test")

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	id := res.InsertedID

	fmt.Printf("%+v\n", id)

	fmt.Println("lol")
	p := NewPomo{}
	p.Start()

}
