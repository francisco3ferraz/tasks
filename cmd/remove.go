package cmd

import (
	"fmt"

	"github.com/francisco3ferraz/tasks/internal/files"
	"github.com/francisco3ferraz/tasks/internal/tasks"
	"github.com/spf13/cobra"
)

var taskId int

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Args != nil {
			fmt.Printf("remove.Run: no task ID provided\n")
			return
		}

		if taskId == 0 {
			fmt.Printf("remove.Run: no task ID provided\n")
			return
		}

		tasks.DeleteTask(taskId)
		files.WriteJSON("tasks.json")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().IntVarP(&taskId, "id", "i", 0, "ID of the task to be removed")
}
