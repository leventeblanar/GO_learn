package main

//  belépési pont

import (
	"fmt"
	"os"
	"strconv"
	"tasktracker/internal/tasks"
)

func usage() {
    fmt.Println("Usage:")
    fmt.Println("  task-tracker add \"title\"")
    fmt.Println("  task-tracker update <id> \"new title\"")
    fmt.Println("  task-tracker delete <id>")
    fmt.Println("  task-tracker mark-in-progress <id>")
    fmt.Println("  task-tracker mark-done <id>")
    fmt.Println("  task-tracker mark-todo <id>")
    fmt.Println("  task-tracker list [todo|in-progress|done]")
}

func main() {
	if len(os.Args) < 2 {
		usage(); return
	}
	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 { usage(); return}
		t, err := tasks.Add(os.Args[2])
		if err != nil { fmt.Println("Error:", err); return }
		fmt.Printf("Task added (ID: %d)\n", t.ID)
	
	case "update":
		if len(os.Args) < 4 { usage(); return }
		id, _ := strconv.Atoi(os.Args[2])
		if err := tasks.Update(id, os.Args[3]); err != nil { fmt.Println("Error:", err) }
	
	case "delete":
		if len(os.Args) < 3 { usage(); return }
		id, _ := strconv.Atoi(os.Args[2])
		if err := tasks.Delete(id); err != nil { fmt.Println("Error:", err) }

	case "mark-in-progress", "mark-done", "mark-todo":
		if len(os.Args) < 3 { usage(); return }
		id, _ := strconv.Atoi(os.Args[2])
		var s tasks.Status
		if cmd == "mark-in-progress" { s = tasks.StatusInProgress}
		if cmd == "mark-done" { s = tasks.StatusDone}
		if cmd == "mark-todo" { s = tasks.StatusToDo}
		if err := tasks.Mark(id, s); err != nil { fmt.Println("Error:", err) }

	case "list":
		var filter *tasks.Status
		if len(os.Args) >= 3 {
			v := tasks.Status(os.Args[2])
			filter = &v
		}
		items, err := tasks.List(filter)
		if err != nil { fmt.Println("Error:", err); return }
		for _, t := range items {
			fmt.Printf("%d [%s] %s\n", t.ID, t.Status, t.Title)
		}

	default:
		usage()
	}
}

