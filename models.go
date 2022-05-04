package main

import "time"

type ToDoItem struct {
	Name        string
	Description string
	DueDate     time.Time
	IsImportant bool
	IsUrgent    bool
}

type ToDoList struct {
	ToDoItems []ToDoItem
}
