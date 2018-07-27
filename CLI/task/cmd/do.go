package cmd

import (
	"fmt"
	"strconv"

	"github.com/MohitArora1/gophercises/CLI/task/repository"
	"github.com/spf13/cobra"
)

var doUsage = `
Usage:
task do [list of task number]
`
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do command is used for doing the task from todo list",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			if id, err := strconv.Atoi(arg); err != nil {
				fmt.Printf("enable to parse the id %v\n", arg)
			} else {
				ids = append(ids, id)
			}

		}
		notExist, err := repository.MarkTaskAsDone(ids)
		if err != nil {
			fmt.Printf("\nnot able to do task\n")
		} else if len(notExist) >= 1 {
			fmt.Printf("\n%v these ids not exist and rest mark as done\n", notExist)
		} else {
			fmt.Println("done tasks")
		}
	},
}

func init() {
	doCmd.SetUsageTemplate(doUsage)
	rootCmd.AddCommand(doCmd)
}
