package main

import (
	"fmt"
	"errors"
)

type ToDo struct {
	id			int
	title		string
	completed	bool
	priority	int
}

type TodoList struct {
	todos		[]ToDo
	nextID		int
}

// AddTodo(title string, priority int) error - új todo (error ha priority nem 1-5)
func (t *TodoList) AddTodo(title string, priority int) error {
	if priority > 5 || priority < 1 {
		return errors.New("a priority 1 és 5 közé kell essen")
	}

	todo := ToDo{
		id: t.nextID,
		title: title,
		completed: false,
		priority: priority,
	}

	t.todos = append(t.todos, todo)
	t.nextID++

	fmt.Printf("Todo hozzáadva: %s\n", todo.title)
	return nil
}
// CompleteTodo(id int) error - befejez egy todot (error ha nincs meg)
func (t *TodoList) CompleteTodo(id int) error {
	for i, todo := range t.todos {
		if todo.id == id {
			t.todos[i].completed = true
			fmt.Println("Sikeres completedre állítás.")
			return nil
		}
	}
	return errors.New("nincs todo ilyen idval")
}

// DeleteTodo(id int) error - töröl egy todot (error ha nincs meg)
func (t *TodoList) DeleteTodo(id int) error {
	for i, todo := range t.todos {
		if todo.id == id {
			t.todos = append(t.todos[:i], t.todos[i+1:]...)
			fmt.Println("Sikeres törlés.")
			return nil
		}
	}
	return errors.New("a termék nem található")
}

// GetPending() []Todo - visszaadja a befejezetlen todokat
func (t TodoList) GetPending() []ToDo {
	var pendingToDo []ToDo
	fmt.Println("Befejezetlen todok:")

	for i, todo := range t.todos {
		if !todo.completed {
			pendingToDo = append(pendingToDo, todo)
			fmt.Printf("%d. %s - %v - %d\n", i+1, todo.title, todo.completed, todo.priority)
		}
	}
	return pendingToDo
}
// GetCompleted() []Todo - visszaadja a befejezett todokat
func (t TodoList) GetCompleted() []ToDo {
	var completedTodo []ToDo
	fmt.Println("Befejezett todok:")

	for i, todo := range t.todos {
		if todo.completed {
			completedTodo = append(completedTodo, todo)
			fmt.Printf("%d. %s - %v - %d\n", i+1, todo.title, todo.completed, todo.priority)
		}
	}
	return completedTodo
}
// ListAll() - kilistázza az összeset (✓ vagy ✗ jelöléssel)
func (t TodoList) ListAll() {
	fmt.Println("Az összes feladat:")
	for i, todo := range t.todos {
		if todo.completed {
			fmt.Printf("%d. %s - %v - %d (✓)\n", i+1, todo.title, todo.completed, todo.priority)
		} else {
			fmt.Printf("%d. %s - %v - %d - (✗)\n", i+1, todo.title, todo.completed, todo.priority)
		}
	}
}

// GetTodoByID(id int) (Todo, error) - ID alapján visszaad egy todot
func (t TodoList) GetTodoByID(id int) (ToDo, error) {
	fmt.Printf("Keresett todo id alapján: %d\n", id)
	for i, todo := range t.todos {
		if id == todo.id {
			fmt.Printf("%d. %s - %v - %d\n", i+1, todo.title, todo.completed, todo.priority)
			return todo, nil
		}
	}
	return ToDo{}, errors.New("nincs ilyen idval rendelkező todo")
}

func main() {
	todolist1 := TodoList {
		nextID: 1,
	}

	todolist1.AddTodo("work", 1)
	todolist1.AddTodo("read", 3)
	todolist1.AddTodo("sleep", 1)
	todolist1.AddTodo("eat", 5)

	if err := todolist1.CompleteTodo(2); err != nil {
		fmt.Println("Hiba a completed-re állítás során: \n", err)
	}

	todolist1.GetPending()

	todolist1.GetCompleted()

	todolist1.ListAll()

	if _, err := todolist1.GetTodoByID(6); err != nil {
		fmt.Println("Hiba a todo Id alapján való keresése során: ", err)
	}
}