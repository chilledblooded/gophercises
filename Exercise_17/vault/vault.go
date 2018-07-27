package vault

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/chilledblooded/gophercises/Exercise_17/cipher"
)

//Vault contains the structure of our vault
type Vault struct {
	encodingKey string
	filePath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

//File function returns Vault object initialised with encoding key and file path.
func File(encodingKey, filePath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filePath:    filePath,
	}
}

//Set method sets the key and value in encrypted format into secret file
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	return v.save()
}

//Get method gets the key and value in plain text format from secret file
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}
	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("NO value found for the key")
	}
	return value, nil
}

func (v *Vault) load() error {
	f, err := os.Open(v.filePath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return nil
	}
	return v.readKeyValues(r)
}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) readKeyValues(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(&v.keyValues)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(w)
}
