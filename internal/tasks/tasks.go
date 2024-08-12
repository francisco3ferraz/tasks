package tasks

import "time"

const DATE_FORMAT = "2006-01-02 15:04:05"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	IsComplete  time.Time `json:"is_complete"`
}

var tasks []*Task

func NewTask(id int, description string) *Task {
	return &Task{ID: id, Description: description, CreatedAt: time.Now(), IsComplete: time.Time{}}
}

func GetTasks() []*Task {
	return tasks
}

func AddTask(task *Task) {
	tasks = append(tasks, task)
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}

func CompleteTask(id int) {
	for _, task := range tasks {
		if task.ID == id && task.IsComplete.IsZero() {
			task.IsComplete = time.Now()
			return
		}
	}
}

func UncompleteTask(id int) {
	for _, task := range tasks {
		if task.ID == id && !task.IsComplete.IsZero() {
			task.IsComplete = time.Time{}
			return
		}
	}
}

func UpdateTasksSlice(newTasks []*Task) {
	tasks = newTasks
}

func GetLastTaskID() int {
	if len(tasks) <= 0 {
		return 0
	}
	return tasks[len(tasks)-1].ID
}
