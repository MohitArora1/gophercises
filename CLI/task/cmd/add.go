package cmd

import (
	"fmt"
	"strings"

	"github.com/MohitArora1/gophercises/CLI/task/repository"
	"github.com/spf13/cobra"
)

var addUsage = `
Usage:
task add <any task>
`
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add command is used to add task into todo",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := repository.InsertIntoDB(task)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Added \"%s\" \n", task)
	},
}

func init() {
	addCmd.SetUsageTemplate(addUsage)
	rootCmd.AddCommand(addCmd)
}
