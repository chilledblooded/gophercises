package main

import (
	"testing"
)

func TestInitApp(t *testing.T) {
	err := initApp()
	if err != nil {
		t.Errorf("Expected Result: No error, Actual Result: Got error : %v", err)
	}
}
func TestMain(t *testing.T) {
	main()
}
