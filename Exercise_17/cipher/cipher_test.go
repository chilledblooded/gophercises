package cipher

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestEncryptWriter(t *testing.T) {
	var w bytes.Buffer
	key := "abc"
	_, err := EncryptWriter(key, &w)
	if err != nil {
		t.Errorf("Expected no err but got err %v", err)
	}
}

func TestDecryptReader(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")

	f, _ := os.Open(fp)
	defer f.Close()
	_, err := DecryptReader("abc", f)
	if err != nil {
		t.Errorf("Expected NO error but got following error : %v ", err)
	}
}
