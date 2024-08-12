package files

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/francisco3ferraz/tasks/internal/tasks"
)

// ReadJSON reads from a JSON file and returns a slice of tasks
func ReadJSON(filepath string) error {
	t := tasks.GetTasks()

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	byteValue, _ := io.ReadAll(f)
	if err := json.Unmarshal(byteValue, &t); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	tasks.UpdateTasksSlice(t)
	return nil
}

// WriteJSON writes a slice of tasks to a JSON file
func WriteJSON(filepath string) error {
	tasks := tasks.GetTasks()

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	w := json.NewEncoder(file)
	w.SetIndent("", "  ")

	if err := w.Encode(tasks); err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}

	ReadJSON(filepath)
	return nil
}
