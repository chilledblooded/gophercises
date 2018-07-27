package cmd

import (
	"fmt"

	"github.com/chilledblooded/gophercises/Exercise_17/vault"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf(" Key : %s  Value : %s", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
