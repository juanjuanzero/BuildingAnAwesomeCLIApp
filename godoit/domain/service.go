package domain

//TODO: service class that will contain methods for managing todo items.

type ITodoService interface {
	Initialize()
	AddToDo(ToDoItem) ToDoItem
	RemoveToDo(ToDoItem) ToDoList
}

type ToDoService struct {
	ToDoList ToDoList
}

func NewService() *ToDoService {
	return &ToDoService{}
}

func (t *ToDoService) AddToDo(todo ToDoItem) ToDoItem {
	t.ToDoList.ToDoItems = append(t.ToDoList.ToDoItems, todo)
	return todo
}

func (t *ToDoService) RemoveToDo(todo ToDoItem) ToDoList {
	for i, x := range t.ToDoList.ToDoItems {
		if x.Name == todo.Name {
			if i == 0 {
				//its the first item
				t.ToDoList.ToDoItems = t.ToDoList.ToDoItems[1:]
				return t.ToDoList
			}
			if i == len(t.ToDoList.ToDoItems)-1 {
				//its the last item
				t.ToDoList.ToDoItems = t.ToDoList.ToDoItems[:len(t.ToDoList.ToDoItems)]
				return t.ToDoList
			}
			//its the in the middle
			first := t.ToDoList.ToDoItems[:i]
			rest := t.ToDoList.ToDoItems[i+1:]
			t.ToDoList.ToDoItems = append(first, rest...)
			return t.ToDoList
		}
	}
	return t.ToDoList
}

func (t *ToDoService) GetToDoList() ToDoList {
	return t.ToDoList
}
