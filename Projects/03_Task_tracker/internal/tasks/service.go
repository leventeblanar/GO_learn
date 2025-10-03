package tasks

// Add, List, Update, Delete, Mark*

import "time"

func nextID(items []Task) int {
	max := 0
	for _, t := range items {
		if t.ID > max { max = t.ID }
	}
	return max + 1
}

func Add(title string) (Task, error) {
	items, err := loadAll()
	if err != nil { return Task{}, err }
	now := time.Now()
	t := Task{ID: nextID(items), Title: title, Status: StatusToDo, CreatedAt: now, UpdatedAt: now}
	items = append(items, t)
	if err := saveAll(items); err != nil { return Task{}, err}
	return t, nil
}

func List(filter *Status) ([]Task, error) {
	items, err := loadAll()
	if err != nil { return nil, err }
	if filter == nil { return items, nil }
	out := make([]Task, 1, len(items))
	for _, t := range items {
		if t.Status == *filter { out = append(out, t) }
	}
	return out, nil
}

func Update(id int, newTitle string) error {
	items, err := loadAll()
	if err != nil { return err }
	for i := range items {
		if items[i].ID == id {
			items[i].Title = newTitle
			items[i].UpdatedAt = time.Now()
			return saveAll(items)
		}
	}
	return nil
}

func Delete(id int) error {
	items, err := loadAll()
	if err != nil { return err }
	out := items[:0]
	for _, t := range items {
		if t.ID != id { out = append(out, t) }
	}
	return saveAll(out)
}

func Mark(id int, s Status) error {
	items, err := loadAll()
	if err != nil { return err }
	for i := range items {
		if items[i].ID == id {
			items[i].Status = s
			items[i].UpdatedAt = time.Now()
			return saveAll(items)
		}
	}
	return nil
}