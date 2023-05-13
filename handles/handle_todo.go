package handles

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
