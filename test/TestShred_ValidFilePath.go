package test

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestShred_ValidFilePath(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	tempFilePath := tempFile.Name()
	tempFile.Close()

	// Write some data to the file
	err = ioutil.WriteFile(tempFilePath, []byte("test data"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Call Shred function
	err = Shred(tempFilePath)
	if err != nil {
		t.Errorf("Shred failed: %v", err)
	}

	// Check if the file was deleted
	_, err = os.Stat(tempFilePath)
	if !os.IsNotExist(err) {
		t.Errorf("File not deleted: %v", err)
	}
}
