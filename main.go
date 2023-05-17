package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// Items that I need to place in non-api

//type PomodoroStore interface {
//	StartPomodoroSession()
//	StartStudySession()
//	StartPomodoroBreak()
//	NavigationMenu(string, string)
//}
//
//type MongoPomodoroStore struct {
//	client *mongo.Client
//	coll   *mongo.Collection
//}
//
//func NewMongoPomodoroStore(client *mongo.Client) *MongoPomodoroStore {
//	return &MongoPomodoroStore{
//		client: client,
//		coll:   client.Database(DBNAME).Collection("pomodoro"),
//	}
//
//}

//// Menu for easy navigation
//func (h *PomodoroHandler) NavigationMenu(nextAction, userInput string) {
//
//	switch userInput {
//	case "y":
//		if nextAction == "study" {
//			h.StartStudySession()
//		} else {
//			h.StartPomodoroBreak()
//		}
//	case "m":
//		h.StartPomodoroSession()
//	//TODO: how can we add functionality here with todo
//	case "t":
//		if nextAction == "add" {
//			//			s.AddTodo()
//		} else {
//			//			s.ListTodos()
//		}
//	case "q":
//		fmt.Println("Quitting...")
//		os.Exit(2)
//	}
//
//	fmt.Println("\nwhere now?")
//}
//
//// StartPomodoroBreak a new break timer
//func (h *PomodoroHandler) StartPomodoroBreak() {
//
//	var duration int
//	var input string
//	fmt.Printf("How long do you want to break for: ")
//	fmt.Scanf("%v", &duration)
//
//	mins := duration * 60
//
//	// start loop over time stated
//	if mins <= 60 {
//		fmt.Printf("pomo break starting for %d minute ..\n\n", mins/60)
//	} else if mins > 60 {
//		fmt.Printf("pomo break starting for %d minutes ..\n\n", mins/60)
//	}
//
//	//TODO: make quit function
//	fmt.Println("q to quit at anytime")
//
//	TimeLoop("Break", mins)
//
//	fmt.Printf("\n\ngo to study (y) yes, (m) menu or (q) quit: ")
//	fmt.Scanf("%v", &input)
//
//	h.NavigationMenu("study", input)
//}
//
//// StartStudySession a new study timer
//func (h *PomodoroHandler) StartStudySession() {
//
//	var duration int
//	var input string
//
//	fmt.Printf("How long do you want to study for: ")
//	fmt.Scanf("%v", &duration)
//
//	mins := duration * 60
//
//	// start loop over time stated
//	if mins <= 60 {
//		fmt.Printf("pomo starting for %d minute ..\n\n", mins/60)
//	} else if mins > 60 {
//		fmt.Printf("pomo starting for %d minutes ..\n\n", mins/60)
//	}
//
//	//TODO: add quit function
//	fmt.Println("q to quit at anytime")
//
//	TimeLoop("Study", mins)
//
//	fmt.Printf("\n\ngo to break (y) yes, (m) menu or (q) quit: ")
//	fmt.Scanf("%v", &input)
//
//	h.NavigationMenu("break", input)
//}
//
//// Start is our init and welcome menu
//func (h *PomodoroHandler) StartPomodoroSession() {
//	//TODO: login user if wanted
//	//TODO: username, password, stored in db
//	input := " "
//	for input != "q" {
//		fmt.Printf("\nDo you want to start a pomo: (y) yes, (t) todo, (q) quit: ")
//
//		_, err := fmt.Scanf("%s", &input)
//		if err != nil {
//			log.Fatal(err)
//			os.Exit(1)
//		}
//
//		h.NavigationMenu("study", input)
//	}
//}

//
//
//func ReadUserInput(prompt string) string {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Print(prompt)
//	text, _ := reader.ReadString('\n')
//	return strings.TrimSpace(text)
//}
//
//// main loop handle case of looping
//func TimeLoop(label string, minutes int) {
//	fmt.Printf("%s starting for %d minutes .. ", label, minutes)
//
//	for i := minutes; i >= 0; i-- {
//		fmt.Printf("\r%s Break Countdown...: %d ", label, i) // \r returns to the start of line
//		time.Sleep(1 * time.Second)
//	}
//
//}
