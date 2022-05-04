package main

//TODO: service class that will contain methods for managing todo items.

var todoList ToDoList

func AddToDo(todo ToDoItem) ToDoItem {
	todoList.ToDoItems = append(todoList.ToDoItems, todo)
	return todo
}

func RemoveToDo(todo ToDoItem) ToDoList {
	for i, t := range todoList.ToDoItems {
		if t.Name == todo.Name {
			if i == 0 {
				//its the first item
				todoList.ToDoItems = todoList.ToDoItems[1:]
				return todoList
			}
			if i == len(todoList.ToDoItems)-1 {
				//its the last item
				todoList.ToDoItems = todoList.ToDoItems[:len(todoList.ToDoItems)]
				return todoList
			}
			//its the in the middle
			first := todoList.ToDoItems[:i]
			rest := todoList.ToDoItems[i+1:]
			todoList.ToDoItems = append(first, rest...)
			return todoList
		}
	}
	return todoList
}

func GetToDoList() ToDoList {
	return todoList
}
