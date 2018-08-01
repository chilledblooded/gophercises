package main

import (
	"path/filepath"

	"github.com/chilledblooded/gophercises/Exercise_7/cmd"
	"github.com/chilledblooded/gophercises/Exercise_7/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	initApp()
}

func initApp() error {
	h, _ := homedir.Dir()
	dbPath := filepath.Join(h, "tasks.db")
	_, err := db.Init(dbPath)
	cmd.RootCmd.Execute()
	return err
}
