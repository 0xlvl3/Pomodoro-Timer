package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// AddTodo will add a todo to a users db
//func (p *NewPomo) AddTodo() {
//	todoStore := db.NewMongoTodoStore(p.client)
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
//	todo, err := todoStore.InsertTodo(context.TODO(), title, description)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("todo added :)", todo)
//
//}

// main loop handle case of looping
func TimeLoop(label string, minutes int) {
	fmt.Printf("%s starting for %d minutes .. ", label, minutes)

	for i := minutes; i >= 0; i-- {
		fmt.Printf("\r%s Break Countdown...: %d ", label, i) // \r returns to the start of line
		time.Sleep(1 * time.Second)
	}

}
