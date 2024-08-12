package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/francisco3ferraz/tasks/internal/tasks"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Args != nil {
			fmt.Fprintf(os.Stderr, "list.Run: no tasks to be displayed\n")
			return
		}

		t := tasks.GetTasks()
		if len(t) <= 0 {
			fmt.Fprintf(os.Stdout, "No tasks to display\n")
			return
		}

		tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
		tabWriter.Write([]byte("ID\tDescription\tCreated At\tIs Complete\n"))
		for _, task := range t {
			completed := "False"
			if !task.IsComplete.IsZero() {
				completed = task.IsComplete.Format(tasks.DATE_FORMAT)
			}

			tabWriter.Write([]byte(fmt.Sprintf("%d\t%s\t%s\t%s\n", task.ID, task.Description, task.CreatedAt.Format(tasks.DATE_FORMAT), completed)))
		}

		tabWriter.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
