package files_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/francisco3ferraz/tasks/internal/files"
	"github.com/francisco3ferraz/tasks/internal/tasks"
)

func setupMockFile(t *testing.T, data interface{}) string {
	t.Helper()
	file, err := os.CreateTemp("", "testfile*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	if data != nil {
		if err := json.NewEncoder(file).Encode(data); err != nil {
			t.Fatalf("failed to write to temp file: %v", err)
		}
	}

	file.Close()
	return file.Name()
}

func TestReadJSON(t *testing.T) {
	mockTasks := []tasks.Task{
		{ID: 1, Description: "Test Task 1"},
		{ID: 2, Description: "Test Task 2"},
	}

	filepath := setupMockFile(t, mockTasks)
	defer os.Remove(filepath)

	if err := files.ReadJSON(filepath); err != nil {
		t.Errorf("ReadJSON failed: %v", err)
	}

	loadedTasks := tasks.GetTasks()
	if len(loadedTasks) != len(mockTasks) {
		t.Errorf("Expected %d tasks, got %d", len(mockTasks), len(loadedTasks))
	}
}

func TestWriteJSON(t *testing.T) {
	filepath := setupMockFile(t, nil)
	defer os.Remove(filepath)

	tasks.UpdateTasksSlice([]*tasks.Task{
		{ID: 1, Description: "Write Test Task"},
	})

	if err := files.WriteJSON(filepath); err != nil {
		t.Errorf("WriteJSON failed: %v", err)
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var writtenTasks []tasks.Task
	if err := json.Unmarshal(content, &writtenTasks); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if len(writtenTasks) != 1 || writtenTasks[0].Description != "Write Test Task" {
		t.Errorf("Unexpected content in JSON file")
	}
}
