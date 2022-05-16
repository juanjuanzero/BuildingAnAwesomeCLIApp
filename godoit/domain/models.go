package domain

type ToDoItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	IsImportant bool   `json:"isImportant"`
	IsUrgent    bool   `json:"isUrgent"`
}

type ToDoList struct {
	ToDoItems []ToDoItem `json:"toDoItems"`
}
