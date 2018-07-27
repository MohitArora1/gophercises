package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI command to handle your todo list",
	Long:  `With task command you can add task to your todo list do the task and view all the task`,
}

// Execute is used to add command under root command
func Execute() error {
	return rootCmd.Execute()
}
