package cmd

import (
	"fmt"
	"os"

	"github.com/francisco3ferraz/tasks/internal/files"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Tasks helps you keep track of you to-dos",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	if err := files.ReadJSON("tasks.json"); err != nil {
		fmt.Fprintf(os.Stderr, "root.init: %v\n", err)
	}
}
