package main

import (
	"path/filepath"

	"github.com/chilledblooded/gophercises/Exercise_7/cmd"
	"github.com/chilledblooded/gophercises/Exercise_7/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	h, _ := homedir.Dir()
	dbPath := filepath.Join(h, "tasks.db")
	db.Init(dbPath)
	cmd.RootCmd.Execute()
}
