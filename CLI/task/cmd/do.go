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
		msg := "Not able to mark task as done"
		if err == nil {
			if len(notExist) >= 1 {
				msg = fmt.Sprintf("\n%v these ids not exist and rest mark as done\n", notExist)
			} else {
				msg = fmt.Sprintln("done tasks")
			}
		}
		fmt.Println(msg)
	},
}

func init() {
	doCmd.SetUsageTemplate(doUsage)
	rootCmd.AddCommand(doCmd)
}
