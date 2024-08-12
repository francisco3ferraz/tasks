package cmd

import (
	"fmt"
	"os"

	"github.com/francisco3ferraz/tasks/internal/files"
	"github.com/francisco3ferraz/tasks/internal/tasks"
	"github.com/spf13/cobra"
)

var task string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Args != nil {
			fmt.Fprintf(os.Stderr, "add.Run: no task provided\n")
			return
		}

		if task == "" {
			fmt.Fprintf(os.Stderr, "add.Run: no task provided\n")
			return
		}

		newTask := tasks.NewTask(tasks.GetLastTaskID()+1, task)
		tasks.AddTask(newTask)

		files.WriteJSON("tasks.json")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&task, "task", "t", "", "Task to be added")
}
