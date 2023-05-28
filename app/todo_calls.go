package app

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

	"github.com/0xlvl3/pomodoro-timer/api/types"
)

// will contain all todo api calls

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
