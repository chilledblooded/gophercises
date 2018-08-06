package cipher

import (
	"bytes"
	"crypto/aes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

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

func TestDecryptReaderNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secrettest.txt")

	f, _ := os.Open(fp)
	defer f.Close()
	_, err := DecryptReader("abc", f)
	if err == nil {
		t.Error("Expected error but got no error")
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

func TestCheckIV(t *testing.T) {
	iv := make([]byte, aes.BlockSize)
	err := checkIV(10, iv, errors.New("test"))
	if err == nil {
		t.Error("Expected error but got no error")
	}
}
func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
