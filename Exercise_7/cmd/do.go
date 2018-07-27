package cmd

import (
	"Gophercises/task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, v := range args {
			id, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error occured: %s", err)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Printf("error occured : %s\n", err)
		}
		for _, i := range ids {
			if i <= 0 || i > len(tasks) {
				fmt.Println("Invalid task number:", i)
				continue
			}
			t := tasks[i-1]
			err := db.DeleteTask(t.ID)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", i, err)
			} else {
				fmt.Printf("Marked your task \"%s\" as completed.\n", t.Task)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
