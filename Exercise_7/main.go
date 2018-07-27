package main

import (
	"Gophercises/task/cmd"
	"Gophercises/task/db"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	h, _ := homedir.Dir()
	dbPath := filepath.Join(h, "tasks.db")
	db.Init(dbPath)
	cmd.RootCmd.Execute()
}
