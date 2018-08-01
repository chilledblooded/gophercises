package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/chilledblooded/gophercises/Exercise_7/db"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestListCommand(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "tasks.db")
	dbc, _ := db.Init(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{""}
	listCmd.Run(listCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "You have the following tasks:")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	dbc.Close()

}

// func TestListCommandNoData(t *testing.T) {
// 	home, _ := homedir.Dir()
// 	DbPath := filepath.Join(home, "taskstest.db")
// 	dbc, _ := db.Init(DbPath)
// 	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 	oldStdout := os.Stdout
// 	a := []string{""}
// 	listCmd.Run(listCmd, a)
// 	file.Seek(0, 0)
// 	content, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		t.Error("error occured while test case : ", err)
// 	}
// 	output := string(content)
// 	fmt.Println(output)
// 	val := strings.Contains(output, "no Task pending")
// 	fmt.Println(val)
// 	assert.Equalf(t, true, val, "they should be equal")
// 	file.Truncate(0)
// 	file.Seek(0, 0)
// 	os.Stdout = oldStdout
// 	fmt.Println(string(content))
// 	file.Close()
// 	dbc.Close()
// }
