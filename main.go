package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/0xlvl3/pomodoro-timer/api/handles"
	"github.com/0xlvl3/pomodoro-timer/api/types"
)

// bring in tea model
// has selections for pre-defined study times 15, 25, 40, 60
// has slected for pre-defined break times 5, 10, 15
// todos have own selection board

func GetUserByEmail() {
	var email string
	fmt.Printf("What is your email: ")
	fmt.Scanln(&email)

	resp, err := http.Get("http://localhost:8080/api/user/" + email)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user types.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Email, user.Username, user.ID)

}

func CreateUser(username, email, password string) {
	user := &types.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/api/user/create", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Trim off the newline at the end
	return input
}

func InsertTodo() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Give your todo a title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title) // Trim off the newline at the end

	fmt.Printf("Describe your todo: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description) // Trim off the newline at the end

	todo := types.Todo{
		Title:       title,
		Description: description,
	}

	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/api/user/todo/add", "application/json", bytes.NewBuffer(jsonTodo))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}

func GetAllTodos(token string) {

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/todo", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Auth-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	var todos []types.Todo
	err = json.Unmarshal(body, &todos)
	if err != nil {
		log.Fatal(err)
	}

	for i, todo := range todos {

		fmt.Printf("%d. Todo \n--- Title: %v \n--- Description: %v \n\n", i, todo.Title, todo.Description)
	}

}

func GetUserByID() {

	resp, err := http.Get("http://localhost:8080/api/test/645df0bd004e4e0b5064590f")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	var user types.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Email, user.Username, user.ID)

}

type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func Login(email, password string) {

	loginUser := handles.AuthParams{
		Email:    email,
		Password: password,
	}

	jsonUser, err := json.Marshal(loginUser)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	fmt.Println(loginResp.Token)

}

func main() {
	for {
		var input int
		fmt.Println("Welcome -- ")
		fmt.Println("Choose from the following options")
		fmt.Println("1 - Log in")
		fmt.Println("2 - Create account")
		fmt.Scanln(&input)

		if input == 1 {
			var email string
			var password string
			fmt.Println("-- Log in --")

			fmt.Printf("Email: ")
			//fmt.Scanln(&email)
			//email := ReadInput()
			fmt.Scanln(&email)

			fmt.Printf("Password: ")
			//password := ReadInput()
			fmt.Scanln(&password)

			Login(email, password)
			continue
		} else if input == 2 {
			var username string
			var email string
			var password string

			fmt.Printf("Choose a username: ")
			fmt.Scanln(&username)

			fmt.Printf("Choose a email: ")
			fmt.Scanln(&email)

			fmt.Printf("Choose a password: ")
			fmt.Scanln(&password)

			CreateUser(username, email, password)
			continue
		} else {
			fmt.Println("Invalid option please enter 1 to log in or 2 to create an account")
			continue
		}
	}
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

//

// was in todo
// AddTodo will add a todo to a users db
//func (h *TodoHandler) AddTodo() {
//
//	fmt.Println("todo")
//
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Printf("Title: ")
//	title, _ := reader.ReadString('\n')
//	title = strings.TrimSpace(title)
//
//	fmt.Printf("Description: ")
//	description, _ := reader.ReadString('\n')
//	description = strings.TrimSpace(description)
//
//	//TODO: add time limit or num of pomos required
//
//	todo, err := h.todoStore.InsertTodo(context.TODO(), title, description)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("todo added :)", todo)
//
//}
//

//

//func StartMenu() {
//	//TODO: get this to work
//
//	fmt.Println("Create account? (c) - create user")
//	fmt.Println("Log in? (l) - log in")
//	fmt.Println("Or continue without one? (p) - proceed")
//	login := ReadUserInput(" ")
//	switch login {
//	case "c":
//		// login
//		username := ReadUserInput("username -- ")
//
//		email := ReadUserInput("email -- ")
//
//		password := ReadUserInput("password -- ")
//
//		user, err := userStore.NewUser(context.TODO(), username, email, password)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(user)
//
//	case "l":
//		fmt.Println("\nLogin")
//	case "p":
//		fmt.Println("\nWelcome Guest")
//	}
//
//	StartPomodoroSession()
//
//}
