package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//Encryt will take key and plain text as arguments and returns the encrytpted text with using key
func Encryt(key, plainText string) (string, error) {
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream, err := encrytStream(key, iv)
	if err != nil {
		return "", err
	}
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))
	return fmt.Sprintf("%x", cipherText), nil
}

//EncryptWriter takes key and io writer as input and decrypt the data and return text in stream writer
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, err := encrytStream(key, iv)
	if err != nil {
		return nil, err
	}
	n, err := w.Write(iv)
	if len(iv) != n || err != nil {
		return nil, errors.New("Unable to write IV into writer")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

//Decrypt function take two arguments key and Hexed cipher and decryt the cipher into plain text
//and return as the string and return error if fails to decrypt
func Decrypt(key, cipherHex string) (string, error) {
	cipherText, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("Cipher too short to decrypt")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream, err := decryptStream(key, iv)
	if err != nil {
		return "", err
	}
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}

//DecryptReader takes key and io reader as input.
//reader will be carring the encrypted data which will be decrypted and returned in stream writer as plain text
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to read the full iv")
	}
	stream, err := decryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func encrytStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newBlockCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newBlockCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBDecrypter(block, iv), nil
}

func newBlockCipher(key string) (cipher.Block, error) {

	hash := md5.New()
	fmt.Fprint(hash, key)
	cipherKey := hash.Sum(nil)
	return aes.NewCipher(cipherKey)
}
