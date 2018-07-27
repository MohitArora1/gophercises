package cmd

import (
	"fmt"

	"github.com/MohitArora1/gophercises/CLI/task/repository"
	"github.com/spf13/cobra"
)

var listUsage = `
Usage:
task list
`

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list command will list all the task in the todo",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := repository.ReadNotCompletedTaskFromDB()
		if err != nil {
			fmt.Print("\nNo task found")
		}
		for i, task := range tasks {
			fmt.Printf("%d. %v\n", i+1, task)
		}
	},
}

func init() {
	listCmd.SetUsageTemplate(listUsage)
	rootCmd.AddCommand(listCmd)

}
