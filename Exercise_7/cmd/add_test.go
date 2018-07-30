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

func TestAddCmd(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "tasks.db")
	dbc, _ := db.Init(DbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"Complete Exercise"}
	addCmd.Run(addCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Added")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
	dbc.Close()
}
