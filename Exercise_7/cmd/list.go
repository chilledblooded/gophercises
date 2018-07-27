package cmd

import (
	"Gophercises/task/db"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Printf("error occured : %s", err)
		}
		if len(tasks) == 0 {
			fmt.Println("you have no Task pending to do...")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Task)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}