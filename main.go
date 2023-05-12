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
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/0xlvl3/pomodoro-timer/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pomo interface {
	PomoStart()
	PomoBreak(int)
}

type NewPomo struct {
	input    string
	duration int
	client   *mongo.Client
}

func (p *NewPomo) readUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// StartPomodoroBreak a new break timer
func (p *NewPomo) StartPomodoroBreak() {
	fmt.Printf("How long do you want to break for: ")
	fmt.Scanf("%v", &p.duration)

	mins := p.duration * 60

	// start loop over time stated
	if mins <= 60 {
		fmt.Printf("pomo break starting for %d minute ..\n\n", mins/60)
	} else if mins > 60 {
		fmt.Printf("pomo break starting for %d minutes ..\n\n", mins/60)
	}

	//TODO: make quit function
	fmt.Println("q to quit at anytime")

	p.TimeLoop("Break", mins)

	fmt.Printf("\n\ngo to study (y) yes, (m) menu or (q) quit: ")
	fmt.Scanf("%v", &p.input)

	p.NavigationMenu("study", p.input)
}

// StartStudySession a new study timer
func (p *NewPomo) StartStudySession() {
	fmt.Printf("How long do you want to study for: ")
	fmt.Scanf("%v", &p.duration)

	mins := p.duration * 60

	// start loop over time stated
	if mins <= 60 {
		fmt.Printf("pomo starting for %d minute ..\n\n", mins/60)
	} else if mins > 60 {
		fmt.Printf("pomo starting for %d minutes ..\n\n", mins/60)
	}

	//TODO: add quit function
	fmt.Println("q to quit at anytime")

	p.TimeLoop("Study", mins)

	fmt.Printf("\n\ngo to break (y) yes, (m) menu or (q) quit: ")
	fmt.Scanf("%v", &p.input)

	p.NavigationMenu("break", p.input)
}

// Start is our init and welcome menu
func (p *NewPomo) StartPomodoroSession() {
	//TODO: login user if wanted
	//TODO: username, password, stored in db

	fmt.Println("Hello user")

	for p.input != "q" {
		fmt.Printf("\nDo you want to start a pomo: (y) yes, (t) todo, (q) quit: ")

		_, err := fmt.Scanf("%s", &p.input)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		p.NavigationMenu("study", p.input)
	}
}

// AddTodo will add a todo to a users db
func (p *NewPomo) AddTodo() {
	todoStore := db.NewMongoTodoStore(p.client)

	fmt.Println("todo")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	//TODO: add time limit or num of pomos required

	todo, err := todoStore.InsertTodo(context.TODO(), title, description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("todo added :)", todo)

}

// Menu for easy navigation
func (p *NewPomo) NavigationMenu(nextAction, userInput string) {
	switch userInput {
	case "y":
		if nextAction == "study" {
			p.StartStudySession()
		} else {
			p.StartPomodoroBreak()
		}
	case "m":
		p.StartPomodoroSession()
	case "t":
		p.AddTodo()
	case "q":
		fmt.Println("Quitting...")
		os.Exit(2)
	}

	fmt.Println("\nwhere now?")
}

// main loop handle case of looping
func (p *NewPomo) TimeLoop(label string, minutes int) {
	fmt.Printf("%s starting for %d minutes .. ", label, minutes)

	for i := minutes; i >= 0; i-- {
		fmt.Printf("\r%s Countdown...: %d ", label, i) // \r returns to the start of line
		time.Sleep(1 * time.Second)
	}

}

// make db for users
// make db for todo

func main() {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("client starting...: ", client)
	p := NewPomo{client: client}
	p.StartPomodoroSession()

}
