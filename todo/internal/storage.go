package internal

import (
	"encoding/json"
	"os"
)

var Todos []Todo

const FileName = "todo.json"

// Load todo from file
func LoadTodos() error {
	file, err := os.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			Todos = []Todo{}
			return nil
		}
		return err
	}

	return json.Unmarshal(file, &Todos)
}

// Save todo to file
func Savetodos() error {
	data, err := json.MarshalIndent(Todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(FileName, data, 0644)
}
