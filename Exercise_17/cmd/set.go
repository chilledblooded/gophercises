package cmd

import (
	"fmt"

	"github.com/chilledblooded/gophercises/Exercise_17/vault"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Printf("Error occured in set %v\n", err)
			return
		}
		fmt.Println("value is added successfully")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
