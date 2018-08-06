package main

import (
	"path/filepath"
	"testing"

	"github.com/chilledblooded/gophercises/Exercise_17/vault"
	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	main()
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := vault.File("abc", fp)
	v.Set("xyz", "testing")
	dashtest.ControlCoverage(m)
}
