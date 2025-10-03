package tasks

// Json load/save (fs only)

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const fileName = "tasks.json"

// visszaadja a fájl elérési útvonalát (./tasks.json)
func filePath() string { return filepath.Join(".", fileName)}


func loadAll() ([]Task, error) {
	p := filePath()

	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(p, []byte("[]"), 0644); err != nil { return nil, err }
	}

	b, err := os.ReadFile(p)
	if err != nil { return nil, err }

	var items []Task
	if err := json.Unmarshal(b, &items); err != nil { return nil, err }
	return items, nil
}

func saveAll(items []Task) error {
	b, err := json.MarshalIndent(items, "", " ")
	if err != nil { return err }
	return os.WriteFile(filePath(), b, 0644)
}