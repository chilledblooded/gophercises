package vault

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestSet(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("abc", fp)
	err := v.Set("xyz", "testing")
	if err != nil {
		t.Errorf("Expected no err but got err %v", err)
	}
}
func TestSetNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("", fp)
	err := v.Set("xyz", "testing")
	if err == nil {
		t.Error("Expected  Error but got suucessful execution")
	}
}

func TestGet(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("abc", fp)
	_, err := v.Get("xyz")
	if err != nil {
		t.Errorf("Expected no err but got err %v", err)
	}
}
func TestGetNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("abc", fp)
	_, err := v.Get("x")
	if err == nil {
		t.Error("Expected Error but got NO error ")
	}
}
