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
		_, err := repository.MarkTaskAsDone(ids)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	doCmd.SetUsageTemplate(doUsage)
	rootCmd.AddCommand(doCmd)
}
