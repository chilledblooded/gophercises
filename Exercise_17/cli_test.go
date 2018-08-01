package main

import (
	"path/filepath"
	"testing"

	"github.com/chilledblooded/gophercises/Exercise_17/vault"
	homedir "github.com/mitchellh/go-homedir"
)

func TestMain(t *testing.T) {
	main()
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := vault.File("abc", fp)
	v.Set("xyz", "testing")
}
