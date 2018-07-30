package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd is a root command of all the task commands
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
