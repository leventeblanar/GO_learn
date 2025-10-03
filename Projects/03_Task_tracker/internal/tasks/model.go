package tasks

//  Task, Status, konstansok

import "time"

type Status string

const (
	StatusToDo			Status = "todo"
	StatusInProgress	Status = "in-progress"
	StatusDone 			Status = "done"
)


type Task struct {
	ID				int 		`json:"id"`
	Title			string		`json:"title"`
	Status			Status		`json:"status"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}