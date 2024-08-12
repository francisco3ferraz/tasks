package tasks_test

import (
	"testing"

	"github.com/francisco3ferraz/tasks/internal/tasks"
)

func FlushTasks() {
	tasks.UpdateTasksSlice([]*tasks.Task{})
}

func TestNewTask(t *testing.T) {
	FlushTasks()

	task := tasks.NewTask(1, "Test task")

	if task.ID != 1 {
		t.Errorf("Expected ID 1, got %d", task.ID)
	}
	if task.Description != "Test task" {
		t.Errorf("Expected description 'Test task', got '%s'", task.Description)
	}
	if task.CreatedAt.IsZero() {
		t.Error("Expected non-zero CreatedAt time")
	}
	if !task.IsComplete.IsZero() {
		t.Error("Expected zero IsComplete time")
	}
}

func TestAddTask(t *testing.T) {
	FlushTasks()

	task := tasks.NewTask(2, "Another task")
	tasks.AddTask(task)

	currentTasks := tasks.GetTasks()
	if len(currentTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(currentTasks))
	}
	if currentTasks[0] != task {
		t.Errorf("Expected task %+v, got %+v", task, currentTasks[0])
	}
}

func TestDeleteTask(t *testing.T) {
	FlushTasks()

	task1 := tasks.NewTask(3, "Task to be deleted")
	task2 := tasks.NewTask(4, "Another task")
	tasks.AddTask(task1)
	tasks.AddTask(task2)

	tasks.DeleteTask(3)

	currentTasks := tasks.GetTasks()
	if len(currentTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(currentTasks))
	}
	if currentTasks[0] != task2 {
		t.Errorf("Expected task %+v, got %+v", task2, currentTasks[0])
	}
}

func TestCompleteTask(t *testing.T) {
	FlushTasks()

	task := tasks.NewTask(5, "Task to be completed")
	tasks.AddTask(task)

	tasks.CompleteTask(5)

	currentTask := tasks.GetTasks()[0]
	if currentTask.IsComplete.IsZero() {
		t.Error("Expected non-zero IsComplete time")
	}
}

func TestUncompleteTask(t *testing.T) {
	FlushTasks()

	task := tasks.NewTask(6, "Task to be uncompleted")
	tasks.CompleteTask(6)
	tasks.AddTask(task)

	tasks.UncompleteTask(6)

	if !task.IsComplete.IsZero() {
		t.Error("Expected zero IsComplete time")
	}
}

func TestUpdateTasksSlice(t *testing.T) {
	FlushTasks()

	task1 := tasks.NewTask(7, "Task 1")
	task2 := tasks.NewTask(8, "Task 2")
	tasks.UpdateTasksSlice([]*tasks.Task{task1, task2})

	currentTasks := tasks.GetTasks()
	if len(currentTasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(currentTasks))
	}
	if currentTasks[0] != task1 || currentTasks[1] != task2 {
		t.Errorf("Expected tasks %+v and %+v, got %+v and %+v", task1, task2, currentTasks[0], currentTasks[1])
	}
}

func TestGetLastTaskID(t *testing.T) {
	FlushTasks()

	task1 := tasks.NewTask(9, "Task 1")
	task2 := tasks.NewTask(10, "Task 2")
	tasks.AddTask(task1)
	tasks.AddTask(task2)

	lastID := tasks.GetLastTaskID()
	if lastID != 10 {
		t.Errorf("Expected last task ID 10, got %d", lastID)
	}
}

func TestGetTasks(t *testing.T) {
	FlushTasks()

	task1 := tasks.NewTask(11, "Task 1")
	task2 := tasks.NewTask(12, "Task 2")
	tasks.AddTask(task1)
	tasks.AddTask(task2)

	currentTasks := tasks.GetTasks()
	if len(currentTasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(currentTasks))
	}
	if currentTasks[0] != task1 || currentTasks[1] != task2 {
		t.Errorf("Expected tasks %+v and %+v, got %+v and %+v", task1, task2, currentTasks[0], currentTasks[1])
	}
}
